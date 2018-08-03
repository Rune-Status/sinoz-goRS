package entity

import (
	"github.com/sinoz/goRS/internal/game_server/game/message"
	"log"
)

type Avatar struct {
	ProcessId int
	Messages  chan interface{}
}

func (avt *Avatar) Poll() {
	avt.pollMessages()
}

func (avt *Avatar) pollMessages() {
	for x := range avt.Messages {
		switch msg := x.(type) {
		case message.ClientFocusChange:
			// TODO
		default:
			log.Printf("Could not find logic block for message %v \n", msg)
		}
	}
}

func (avt *Avatar) Update() {
	// TODO
}

func (avt *Avatar) Reset() {
	// TODO
}