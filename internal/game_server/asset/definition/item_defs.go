package definition

import (
	"io/ioutil"
	"encoding/json"
)

type ItemDefinition struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	BagOptions        []string `json:"bagOptions"`
	FloorOptions      []string `json:"floorOptions"`
	BankPlaceholderID int      `json:"bankPlaceholderId"`
	Tradeable         bool     `json:"tradeable"`
	Stackable         bool     `json:"stackable"`
	Weight            float64  `json:"weight"`
	Examine           string   `json:"examine"`
}

func LoadItemDefsFromFile(path string) ([]ItemDefinition, error) {
	jsonData, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		return nil, loadErr
	}

	var definitions []ItemDefinition

	unmarshallErr := json.Unmarshal(jsonData, &definitions)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return definitions, nil
}