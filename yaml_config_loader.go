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

	for authName, _ := range roleDraft {
		if authName == "name" {
			continue
		}

		users = append(users, getUsers(roleDraft, authName)...)
		groups = append(groups, getGroups(roleDraft, authName)...)
	}
	return models.AuthRule{users, groups}
}

func getUsers(roleDraft map[string]interface{}, authName string) []string {
	return getValues(roleDraft, authName, usersKeys)
}

func getGroups(roleDraft map[string]interface{}, authName string) []string {
	return getValues(roleDraft, authName, groupsKeys)
}

func getValues(roleDraft map[string]interface{}, authName string, keysTable map[string][]string) []string {
	rule := roleDraft[authName].(map[interface{}]interface{})

	values := make([]string, 0)
	for _, key := range keysTable[authName] {
		if list := rule[key]; list != nil {
			for _, name := range list.([]interface{}) {
				prefixedName := authName + ":" + name.(string)
				values = append(values, prefixedName)
			}
		}
	}
	return values
}
