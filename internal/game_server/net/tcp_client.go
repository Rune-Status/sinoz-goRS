package net

import (
	"net"
	"bufio"
	"log"
	"sync"
	"github.com/sinoz/telos/internal/game_server/login"
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

receiveData:
	for {
		opcode, err := client.reader.ReadByte()
		if err != nil {
			break receiveData
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
			in.Fill(client.reader, 10)

			in.ReadInt8() // magic value
			in.ReadInt16() // client revision

			response := NewPacket(3)
			response.WriteInt8(login.LoginSuccess)
			response.WriteInt8(0)
			response.WriteInt8(0)
			response.WriteAndFlush(client.writer)

			stage = IngameStage

		case IngameStage:
			// TODO
		}

	}
}

func (client *TcpClient) connectionTerminated() {
	close(client.downstream)
	close(client.upstream)
}

func (client *TcpClient) Write() {
	for {
		for message := range client.downstream {
			switch msg := message.(type) {
			default:
				log.Fatalf("Could not find implementation for message %v \n", msg)
			}
		}
	}
}

func (client *TcpClient) Flush() {
	// TODO
}
