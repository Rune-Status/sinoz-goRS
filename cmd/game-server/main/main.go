package main

import (
	"github.com/sinoz/goRS/internal/game_server/net"
	"github.com/sinoz/goRS/internal/game_server/login"
	"github.com/sinoz/goRS/internal/game_server/game"
)

func main() {
	gameService := game.NewService()
	loginService := login.NewService(gameService)

	tcpServer := net.NewTcpServer(43594, loginService)
	tcpServer.StartListening()
}