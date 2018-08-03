package definition

import (
	"io/ioutil"
	"encoding/json"
)

type BitVariableDefinition struct {
	Id          int `json:"bitVariableId"`
	VariableId  int `json:"variableId"`
	LeastSigBit int `json:"leastSigBit"`
	MostSigBit  int `json:"mostSigBit"`
}

func LoadBitVariableDefsFromFile(path string) ([]BitVariableDefinition, error) {
	jsonData, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		return nil, loadErr
	}

	var definitions []BitVariableDefinition

	unmarshallErr := json.Unmarshal(jsonData, &definitions)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return definitions, nil
}
