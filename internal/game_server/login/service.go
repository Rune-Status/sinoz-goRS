package login

import (
	"math/rand"
	"github.com/sinoz/goRS/internal/game_server/game"
)

type Service struct {
	gameService *game.Service
}

func NewService(gameService *game.Service) *Service {
	return &Service{gameService: gameService}
}

func (service *Service) Start() {
	// TODO
}

func (service *Service) Stop() {
	// TODO
}

func (service *Service) GenerateSessionKey() int64 {
	return rand.Int63()
}
