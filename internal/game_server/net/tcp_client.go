package net

import (
	"net"
	"bufio"
	"log"
	"sync"
	"github.com/sinoz/telos/internal/game_server/login"
	"github.com/sinoz/telos/internal/game_server/message"
)

const (
	HandshakeStage = 0
	LoginStage     = 1
	IngameStage    = 2
)

type TcpClient struct {
	connection     net.Conn
	reader         *bufio.Reader
	writer         *bufio.Writer
	upstreamPool   sync.Pool
	downstreamPool sync.Pool
	upstream       chan UpstreamMessage
	downstream     chan DownstreamMessage
	loginService   *login.Service
}

type UpstreamMessage interface{}
type DownstreamMessage interface{}

func NewTcpClient(connnection net.Conn, upstreamPool, downstreamPool sync.Pool, login *login.Service) *TcpClient {
	return &TcpClient{
		connection:     connnection,
		reader:         bufio.NewReader(connnection),
		writer:         bufio.NewWriter(connnection),
		upstreamPool:   upstreamPool,
		downstreamPool: downstreamPool,
		upstream:       make(chan UpstreamMessage, 64),
		downstream:     make(chan DownstreamMessage, 256),
		loginService:   login,
	}
}

func (client *TcpClient) Read() {
	defer client.connectionTerminated()

	in := client.upstreamPool.Get().(*Packet)

	stage := HandshakeStage

connectionLoop:
	for {
		opcode, err := client.reader.ReadByte()
		if err != nil {
			break connectionLoop
		}

		switch stage {
		case HandshakeStage:
			if opcode == 14 {
				in.Fill(client.reader, 1)

				in.ReadInt8() // partial name hash

				sessionKey := client.loginService.GenerateSessionKey()

				response := NewPacket(17)
				response.WriteInt8(login.MayProceed)
				response.WriteInt64(0)
				response.WriteInt64(sessionKey)
				response.WriteAndFlush(client.writer)

				stage = LoginStage
			}

		case LoginStage:
			// wait for the size notation to come in
			in.Fill(client.reader, 1)

			// read the size and then wait for all of the remaining bytes to come in
			payloadSize := in.ReadInt8()
			in.Fill(client.reader, int(payloadSize))

			// and then read the login payload

			in.ReadInt8() // magic value
			in.ReadInt16() // client version

			in.ReadBool() // low mem version

			archiveCRCs := make([]int, 9)
			for i := 0; i < len(archiveCRCs); i++ {
				archiveCRCs[i] = int(in.ReadInt32())
			}

			in.ReadInt8() // rsa block size

			rsaBlockId := in.ReadInt8()
			if rsaBlockId != 10 {
				client.sendLoginFailure(login.LoginServerRejected)
				break connectionLoop
			}

			seeds := make([]int, 4)
			for i := 0; i < len(seeds); i++ {
				seeds[i] = int(in.ReadInt32())
			}

			in.ReadInt32() // uid

			in.ReadCString() // username
			in.ReadCString() // password

			client.downstream <- message.SuccesfulLogin{Rank: 2, Flagged: false}

			stage = IngameStage

		case IngameStage:
			// TODO
		}
	}
}

func (client *TcpClient) sendLoginFailure(responseCode int) {
	response := NewPacket(1)
	response.WriteInt8(int8(responseCode))
	response.WriteAndFlush(client.writer)
}

func (client *TcpClient) connectionTerminated() {
	close(client.downstream)
	close(client.upstream)
}

func (client *TcpClient) Write() {
	for {
		for downstreamMessage := range client.downstream {
			switch msg := downstreamMessage.(type) {
			case message.SuccesfulLogin:
				response := NewPacket(3)

				response.WriteInt8(login.LoginSuccess)
				response.WriteInt8(int8(msg.Rank))
				response.WriteBool(msg.Flagged)

				response.WriteAndFlush(client.writer)

			default:
				log.Fatalf("Could not find implementation for downstreamMessage %v \n", msg)
			}
		}
	}
}

func (client *TcpClient) Flush() {
	// TODO
}
