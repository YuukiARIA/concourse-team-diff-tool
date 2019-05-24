package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/YuukiARIA/concourse-team-diff-tool/models"
	"github.com/YuukiARIA/concourse-team-diff-tool/stringSet"
)

var (
	colorOfCreated  = color.New(color.FgGreen)
	colorOfDeleted  = color.New(color.FgRed)
	colorOfRetained = color.New(color.FgWhite)
)

func Compare(oldTeam, newTeam models.Team) {
	for roleName, oldRule := range oldTeam.Auth {
		newRule, exists := newTeam.Auth[roleName]
		if exists {
			fmt.Println("role: " + roleName)
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
	fmt.Println("users:")
	compareIds(oldRule.Users, newRule.Users)
	fmt.Println("groups:")
	compareIds(oldRule.Groups, newRule.Groups)
}

func compareIds(oldIds, newIds []string) {
	oldIdsSet, newIdsSet := stringSet.New(oldIds...), stringSet.New(newIds...)
	deletedIdsSet, createdIdsSet, retainedIdsSet := stringSet.Partition(oldIdsSet, newIdsSet)

	showAsCreated(createdIdsSet.Array())
	showAsDeleted(deletedIdsSet.Array())
	showAsRetained(retainedIdsSet.Array())
}

func showAsCreated(values []string) {
	showWithColor(values, "+", colorOfCreated)
}

func showAsDeleted(values []string) {
	showWithColor(values, "-", colorOfDeleted)
}

func showAsRetained(values []string) {
	showWithColor(values, " ", colorOfRetained)
}

func showWithColor(values []string, prefix string, color *color.Color) {
	for _, value := range values {
		color.Printf("%s %s\n", prefix, value)
	}
}
