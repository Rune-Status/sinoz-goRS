package definition

import "testing"

func TestLoadInventoryDefsFromFile(t *testing.T) {
	definitions, err := LoadInventoryDefsFromFile("assets/definition/inventories.json")
	if err != nil {
		t.Error(err)
		return
	}

	inv := definitions[93]
	if inv.Capacity != 28 {
		t.Fatalf("capacity of inventory 93 did not match 28 but instead equals %v \n", inv.Capacity)
	}
}
