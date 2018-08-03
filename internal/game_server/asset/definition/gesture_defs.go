package definition

import (
	"io/ioutil"
	"encoding/json"
)

type GestureDefinition struct {
	ID       int `json:"id"`
	Duration int `json:"durationInMs"`
}

func LoadGestureDefsFromFile(path string) ([]GestureDefinition, error) {
	jsonData, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		return nil, loadErr
	}

	var definitions []GestureDefinition

	unmarshallErr := json.Unmarshal(jsonData, &definitions)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return definitions, nil
}
