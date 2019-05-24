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
		"local":           {"users"},
		"github":          {"users"},
		"gitlab":          {"users"},
		"bitbucket-cloud": {"users"},
		"cf":              {"users"},
		"ldap":            {"users"},
		"oidc":            {"users"},
		"oauth":           {"users"},
	}
	groupsKeys = map[string][]string{
		"local":           {},
		"github":          {"teams", "orgs"},
		"gitlab":          {"teams", "orgs"},
		"bitbucket-cloud": {"teams"},
		"cf":              {"orgs", "spaces"},
		"ldap":            {"groups"},
		"oidc":            {"groups"},
		"oauth":           {"groups"},
	}
)

func LoadYAML(yamlData []byte) models.Team {
	teamDraft := teamDraft{}
	yaml.Unmarshal(yamlData, &teamDraft)

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

	for authName, value := range roleDraft {
		rule, ok := value.(map[interface{}]interface{})
		if !ok {
			continue
		}
		users = append(users, getUsers(rule, authName)...)
		groups = append(groups, getGroups(rule, authName)...)
	}
	return models.AuthRule{users, groups}
}

func getUsers(rule map[interface{}]interface{}, authName string) []string {
	return getValues(rule, authName, usersKeys)
}

func getGroups(rule map[interface{}]interface{}, authName string) []string {
	return getValues(rule, authName, groupsKeys)
}

func getValues(rule map[interface{}]interface{}, authName string, keysTable map[string][]string) []string {
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
