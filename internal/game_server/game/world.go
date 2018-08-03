package game

import (
	"github.com/sinoz/goRS/internal/game_server/game/collection"
	"github.com/sinoz/goRS/internal/game_server/game/entity"
)

type World struct {
	avatars    *collection.Index
	npcs       *collection.Index
	avatarPIDs chan int
	npcPIDs    chan int
}

const (
	AvatarLimit = 1 << 11
	NpcLimit    = 1 << 15
)

func NewWorld() *World {
	world := &World{
		avatars:    collection.NewIndex(1, AvatarLimit),
		npcs:       collection.NewIndex(1, NpcLimit),
		avatarPIDs: make(chan int, AvatarLimit),
		npcPIDs:    make(chan int, NpcLimit),
	}

	for i := 1; i < AvatarLimit; i++ {
		world.avatarPIDs <- i
	}

	for i := 1; i < NpcLimit; i++ {
		world.npcPIDs <- i
	}

	return world
}

func (world *World) RegisterAvatar(avatar *entity.Avatar) error {
	return world.avatars.Set(avatar.ProcessId, avatar)
}

func (world *World) UnregisterAvatar(avatar *entity.Avatar) {
	world.avatars.Set(avatar.ProcessId, nil)
	world.avatarPIDs <- avatar.ProcessId
}

func (world *World) RegisterNpc(npc *entity.Npc) error {
	return world.npcs.Set(npc.ProcessId, npc)
}

func (world *World) UnregisterNpc(npc *entity.Npc) {
	world.npcs.Set(npc.ProcessId, nil)
	world.npcPIDs <- npc.ProcessId
}

func (world *World) GetNextAvatarPid() int {
	return <-world.avatarPIDs
}

func (world *World) GetNextNpcPid() int {
	return <-world.npcPIDs
}
