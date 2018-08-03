package game

import (
	"time"
	"log"
)

const (
	TickRate = 600 * time.Millisecond
)

type Service struct {
	Assets Assets
	World  *World
}

func NewService(assets Assets) *Service {
	return &Service{
		Assets: assets,
		World: NewWorld(),
	}
}

func (service *Service) Start() {
	go func() {
		for {
			start := time.Now()

			err := service.tick()
			if err != nil {
				log.Println(err)
			}

			elapsed := time.Since(start)
			timeRemaining := TickRate - elapsed
			if timeRemaining > 0 {
				time.Sleep(timeRemaining)
			} else if timeRemaining < 0 {
				log.Printf("[Warning]: Game tick took longer than %v", TickRate)
			}
		}
	}()
}

func (service *Service) tick() error {
	return nil
}

func (service *Service) Stop() {
	// TODO
}