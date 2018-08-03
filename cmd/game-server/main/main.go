package main

import (
	"github.com/sinoz/goRS/internal/game_server/net"
	"github.com/sinoz/goRS/internal/game_server/login"
	"github.com/sinoz/goRS/internal/game_server/game"
	"github.com/sinoz/goRS/internal/game_server/asset/definition"
	"log"
)

func main() {
	assets, err := loadAssets()
	if err != nil {
		log.Fatal(err)
	}

	logLoadedAssets(*assets)

	gameService := game.NewService(*assets)
	loginService := login.NewService(gameService)

	tcpServer := net.NewTcpServer(43594, loginService)
	tcpServer.StartListening()
}

func loadAssets() (*game.Assets, error) {
	items, err := definition.LoadItemDefsFromFile("assets/definition/items.json")
	if err != nil {
		return nil, err
	}

	npcs, err := definition.LoadNpcDefsFromFile("assets/definition/npcs.json")
	if err != nil {
		return nil, err
	}

	return &game.Assets{Items: items, Npcs: npcs}, nil
}

func logLoadedAssets(assets game.Assets) {
	log.Printf("Loaded %v item definitions \n", len(assets.Items))
	log.Printf("Loaded %v npc definitions \n", len(assets.Npcs))
}