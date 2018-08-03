package game

import (
	"testing"
	"github.com/sinoz/goRS/internal/game_server/game/entity"
)

func TestWorld_GetNextAvatarPid(t *testing.T) {
	world := NewWorld()

	pid := world.GetNextAvatarPid()
	if pid != 1 {
		t.Fatalf("avatar pid expected to be 1 but was actually %v", pid)
	}
}

func TestWorld_GetNextNpcPid(t *testing.T) {
	world := NewWorld()

	pid := world.GetNextNpcPid()
	if pid != 1 {
		t.Fatalf("npc pid expected to be 1 but was actually %v", pid)
	}
}

func TestWorld_UnregisterAvatar(t *testing.T) {
	world := NewWorld()

	pid := world.GetNextAvatarPid()
	avatar := &entity.Avatar{ProcessId: pid}

	world.UnregisterAvatar(avatar)
	nextPid := world.GetNextAvatarPid()
	if nextPid != 2 {
		t.Fatalf("next avatar pid expected to be 2 but was actually %v", nextPid)
	}
}

func TestWorld_UnregisterNpc(t *testing.T) {
	world := NewWorld()

	pid := world.GetNextNpcPid()
	npc := &entity.Npc{ProcessId: pid}

	world.UnregisterNpc(npc)
	nextPid := world.GetNextNpcPid()
	if nextPid != 2 {
		t.Fatalf("next npc pid expected to be 2 but was actually %v", nextPid)
	}
}
