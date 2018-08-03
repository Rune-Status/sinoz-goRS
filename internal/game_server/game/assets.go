package game

import "github.com/sinoz/goRS/internal/game_server/asset/definition"

type Assets struct {
	Items       []definition.ItemDefinition
	Npcs        []definition.NpcDefinition
	Objects     []definition.ObjectDefinition
	Inventories []definition.InventoryDefinition
	Gestures    []definition.GestureDefinition
}
