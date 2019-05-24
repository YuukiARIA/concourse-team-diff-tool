package main

import (
	"fmt"

	"github.com/YuukiARIA/concourse-team-diff-tool/models"

	yaml "gopkg.in/yaml.v2"
)

// Example
//
// roles:
//   - name: owner
//     local:
//       users:
//         - user1
//         - user2
//     github:
//       users:
//         - user3
//       teams:
//         - some-org:some-team
type teamDraft struct {
	Roles []map[string]interface{} `yaml:"roles"`
}

var (
	usersKeys = map[string][]string{
		"local":  {"users"},
		"github": {"users"},
	}
	groupsKeys = map[string][]string{
		"local":  {},
		"github": {"teams", "orgs"},
	}
)

func LoadYAML(yamlData []byte) models.Team {
	teamDraft := teamDraft{}
	yaml.Unmarshal(yamlData, &teamDraft)
	fmt.Printf("TEAM DRAFT\n%#v\n", teamDraft)

	team := models.NewEmpty()
	for _, role := range teamDraft.Roles {
		roleName, ok := role["name"].(string)
		if !ok {
			continue
		}
		team.Auth[roleName] = convertToAuthRule(role)
	}
	return team
}

func convertToAuthRule(roleDraft map[string]interface{}) models.AuthRule {
	users, groups := make([]string, 0), make([]string, 0)
	return models.AuthRule{users, groups}
}
