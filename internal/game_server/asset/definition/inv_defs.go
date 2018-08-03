package definition

import (
	"io/ioutil"
	"encoding/json"
)

type InventoryDefinition struct {
	ID       int `json:"id"`
	Capacity int `json:"capacity"`
}

func LoadInventoryDefsFromFile(path string) ([]InventoryDefinition, error) {
	jsonData, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		return nil, loadErr
	}

	var definitions []InventoryDefinition

	unmarshallErr := json.Unmarshal(jsonData, &definitions)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return definitions, nil
}
