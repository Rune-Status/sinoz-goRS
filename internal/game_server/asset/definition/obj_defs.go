package definition

import (
	"io/ioutil"
	"encoding/json"
)

type ObjectDefinition struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Width        int      `json:"width"`
	Length       int      `json:"length"`
	ClipType     int      `json:"clipType"`
	Impenetrable bool     `json:"impenetrable"`
	MapAreaID    int      `json:"mapAreaId"`
	IsSolid      bool     `json:"isSolid"`
	Options      []string `json:"options"`
	MotionID     int      `json:"motionId"`
	VarpID       int      `json:"varpId"`
	VarbitID     int      `json:"varbitId"`
}

func LoadObjectDefsFromFile(path string) ([]ObjectDefinition, error) {
	jsonData, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		return nil, loadErr
	}

	var definitions []ObjectDefinition

	unmarshallErr := json.Unmarshal(jsonData, &definitions)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return definitions, nil
}
