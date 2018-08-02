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

type flush struct{}

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

			in.ReadInt8()  // magic value
			in.ReadInt16() // client version

			in.ReadBool() // low mem version

			archiveCRCs := make([]int, 9)
			for i := 0; i < len(archiveCRCs); i++ {
				archiveCRCs[i] = int(in.ReadInt32())
			}

			in.ReadInt8() // rsa block size

			rsaBlockId := in.ReadInt8()
			if rsaBlockId != 10 {
				client.Enqueue(message.FailedLogin{ResponseCode: login.LoginServerRejected})
				client.Flush()

				break connectionLoop
			}

			seeds := make([]int, 4)
			for i := 0; i < len(seeds); i++ {
				seeds[i] = int(in.ReadInt32())
			}

			in.ReadInt32() // uid

			in.ReadCString() // username
			in.ReadCString() // password

			client.Enqueue(message.SuccesfulLogin{Rank: 2, Flagged: false})
			client.Enqueue(message.Details{ProcessId: 1, Members: true})

			client.Flush()

			stage = IngameStage

		case IngameStage:
			if opcode == 3 {
				in.Fill(client.reader, 1)
				in.ReadBool()
			} else {
				log.Fatalf("Could not find decoder block for message %v", opcode)
			}
		}
	}

	// and put the upstream buffer back into its pool so it can be reused
	client.upstreamPool.Put(in)
}

func (client *TcpClient) Write() {
	out := client.downstreamPool.Get().(*Packet)

	for downstreamMessage := range client.downstream {
		switch msg := downstreamMessage.(type) {
		case message.SuccesfulLogin:
			out.WriteInt8(login.LoginSuccess)
			out.WriteInt8(msg.Rank)
			out.WriteBool(msg.Flagged)

		case message.FailedLogin:
			out.WriteInt16(msg.ResponseCode)

		case message.Details:
			out.WriteInt8(249)
			out.WriteBool(msg.Members)
			out.WriteInt16(msg.ProcessId)

		case message.Logout:
			out.WriteInt8(109)

		case message.SkillUpdate:
			out.WriteInt8(134)
			out.WriteInt8(msg.Id)
			out.WriteInt32(int(msg.Experience))
			out.WriteInt8(msg.Level)

		case flush:
			out.WriteAndFlush(client.writer)

		default:
			log.Fatalf("Could not find implementation for downstream message %v \n", msg)
		}
	}

	// and put the downstream buffer back into its pool so it can be reused
	client.downstreamPool.Put(out)
}

func (client *TcpClient) connectionTerminated() {
	close(client.downstream)
	close(client.upstream)
}

func (client *TcpClient) Enqueue(msg DownstreamMessage) {
	client.downstream <- msg
}

func (client *TcpClient) Flush() {
	client.downstream <- flush{}
}
