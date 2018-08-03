package definition

import "testing"

func TestLoadNpcDefsFromFile(t *testing.T) {
	definitions, err := LoadItemDefsFromFile("assets/npcs.json")
	if err != nil {
		t.Error(err)
		return
	}

	lumdo := definitions[1453]
	if lumdo.Name != "Lumdo" {
		t.Fatalf("name of npc 1453 did not match 'Lumdo' but instead equals %v \n", lumdo.Name)
	}
}
