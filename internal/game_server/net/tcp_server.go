package net

import (
	"net"
	"strconv"
	"log"
	"sync"
	"github.com/sinoz/goRS/internal/game_server/login"
)

type TcpServer struct {
	Port         int
	Listener     net.Listener
	loginService *login.Service
}

func NewTcpServer(port int, login *login.Service) *TcpServer {
	return &TcpServer{
		Port:         port,
		loginService: login,
	}
}

func (server *TcpServer) StartListening() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Local channel bound at %v \n", server.Port)

	defer func() {
		listener.Close()

		log.Println("Local channel unbound")
	}()

	upstreamPool := sync.Pool{New: func() interface{} { return NewPacket(1024) }}
	downstreamPool := sync.Pool{New: func() interface{} { return NewPacket(16384) }}

	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}

		client := NewTcpClient(connection, upstreamPool, downstreamPool, server.loginService)

		go client.Read()
		go client.Write()
	}
}

func (server *TcpServer) StopListening() {
	if server.Listener != nil {
		server.Listener.Close()
	}
}
