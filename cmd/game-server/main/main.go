package main

import (
	"github.com/sinoz/goRS/internal/game_server/net"
	"github.com/sinoz/goRS/internal/game_server/login"
)

func main() {
	loginService := &login.Service{}

	tcpServer := net.NewTcpServer(43594, loginService)
	tcpServer.StartListening()
}