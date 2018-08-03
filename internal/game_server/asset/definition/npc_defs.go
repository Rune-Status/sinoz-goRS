package definition

import (
	"io/ioutil"
	"encoding/json"
)

type NpcDefinition struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Size             int           `json:"size"`
	Stance           int           `json:"stance"`
	Walk             int           `json:"walk"`
	Options          []interface{} `json:"options"`
	VisibleOnMinimap bool          `json:"visibleOnMinimap"`
	Clickable        bool          `json:"clickable"`
	VarbitID         int           `json:"varbitId"`
	VarpID           int           `json:"varpId"`
	CombatStats      []int         `json:"combatStats"`
	CombatBonuses    []int         `json:"combatBonuses"`
	AttackSpeed      int           `json:"attackSpeed"`
	SlayerExpGain    float64       `json:"slayerExpGain"`
	ImmuneToPoison   bool          `json:"immuneToPoison"`
	ImmuneToVenom    bool          `json:"immuneToVenom"`
	Examine          string        `json:"examine"`
}

func LoadNpcDefsFromFile(path string) ([]NpcDefinition, error) {
	jsonData, loadErr := ioutil.ReadFile(path)
	if loadErr != nil {
		return nil, loadErr
	}

	var definitions []NpcDefinition

	unmarshallErr := json.Unmarshal(jsonData, &definitions)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return definitions, nil
}