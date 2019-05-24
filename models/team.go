package models

import (
	"encoding/json"
)

type Team struct {
	ID   int                 `json:"id"`
	Name string              `json:"name"`
	Auth map[string]AuthRule `json:"auth"`
}

type AuthRule struct {
	Users  []string `json:"users"`
	Groups []string `json:"groups"`
}

// Helper for construction by unmarshaling json bytes.
func NewFromJSON(jsonData []byte) Team {
	team := Team{}
	json.Unmarshal(jsonData, &team)
	return team
}
