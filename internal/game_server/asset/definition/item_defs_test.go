package definition

import "testing"

func TestLoadItemDefsFromFile(t *testing.T) {
	definitions, err := LoadItemDefsFromFile("assets/items.json")
	if err != nil {
		t.Error(err)
		return
	}

	abbyWhip := definitions[4151]
	if abbyWhip.Name != "Abyssal whip" {
		t.Fatalf("name of item 4151 did not match 'Abyssal whip' but instead equals %v \n", abbyWhip.Name)
	}
}
