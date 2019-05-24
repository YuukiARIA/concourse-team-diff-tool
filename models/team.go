package models

import (
	"encoding/json"
)

type Team struct {
	ID   int                  `json:"id"`
	Name string               `json:"name"`
	Auth map[string]*AuthRule `json:"auth"`
}

type AuthRule struct {
	RoleName string
	Users    []string `json:"users"`
	Groups   []string `json:"groups"`
}

func NewEmpty() Team {
	return Team{Auth: map[string]*AuthRule{}}
}

// Helper for construction by unmarshaling json bytes.
func NewFromJSON(jsonData []byte) Team {
	team := Team{}
	json.Unmarshal(jsonData, &team)

	for roleName, authRule := range team.Auth {
		authRule.RoleName = roleName
	}

	return team
}
