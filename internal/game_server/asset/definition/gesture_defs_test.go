package definition

import "testing"

func TestLoadGestureDefsFromFile(t *testing.T) {
	definitions, err := LoadGestureDefsFromFile("assets/gestures.json")
	if err != nil {
		t.Error(err)
		return
	}

	gesture := definitions[1282]
	if gesture.Duration != 1320 {
		t.Fatalf("duration of gesture 1282 did not match 1320 but instead equals %v \n", gesture.Duration)
	}
}
