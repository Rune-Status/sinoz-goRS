package entity

import "log"

type Npc struct {
	ProcessId int
	Messages  chan interface{}
}

func (npc *Npc) Poll() {
	npc.pollMessages()
}

func (npc *Npc) pollMessages() {
	for x := range npc.Messages {
		switch msg := x.(type) {
		default:
			log.Printf("Could not find logic block for message %v \n", msg)
		}
	}
}

func (npc *Npc) Update() {
	// TODO
}

func (npc *Npc) Reset() {
	// TODO
}
