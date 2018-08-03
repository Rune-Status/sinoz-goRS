package definition

import "testing"

func TestLoadBitVariableDefsFromFile(t *testing.T) {
	definitions, err := LoadBitVariableDefsFromFile("assets/bit_variables.json")
	if err != nil {
		t.Error(err)
		return
	}

	bitVariable := definitions[552]
	if bitVariable.VariableId != 468 {
		t.Fatalf("variable id of bit variable 552 did not match 468 but instead equals %v \n", bitVariable.VariableId)
	}

	if bitVariable.LeastSigBit != 25 {
		t.Fatalf("least significant bit of bit variable 552 did not match 25 but instead equals %v \n", bitVariable.LeastSigBit)
	}

	if bitVariable.MostSigBit != 25 {
		t.Fatalf("least significant bit of bit variable 552 did not match 25 but instead equals %v \n", bitVariable.LeastSigBit)
	}
}
