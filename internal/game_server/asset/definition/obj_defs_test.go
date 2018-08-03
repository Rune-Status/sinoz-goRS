package definition

import "testing"

func TestLoadObjectDefsFromFile(t *testing.T) {
	definitions, err := LoadObjectDefsFromFile("assets/definition/objects.json")
	if err != nil {
		t.Error(err)
		return
	}

	obj := definitions[33108]
	if obj.Name != "Yanille Watchtower Portal" {
		t.Fatalf("name of object 33108 did not match 'Yanille Watchtower Portal' but instead equals %v \n", obj.Name)
	}
}
