package main

import (
	"fmt"

	"github.com/YuukiARIA/concourse-team-diff-tool/models"
)

func Compare(oldTeam, newTeam models.Team) {
	for roleName, oldRule := range oldTeam.Auth {
		newRule, exists := newTeam.Auth[roleName]
		if exists {
			compareRule(oldRule, newRule)
		} else {
			fmt.Println("role definition deleted: " + roleName)
		}
	}
	for roleName, _ := range newTeam.Auth {
		_, exists := oldTeam.Auth[roleName]
		if !exists {
			fmt.Println("role newly defined: " + roleName)
		}
	}
}

func compareRule(oldRule, newRule models.AuthRule) {
	// TODO
}
