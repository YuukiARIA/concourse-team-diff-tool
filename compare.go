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
	colorOfStrong   = color.New(color.FgHiWhite).Add(color.Bold)
)

type compareResult []compareRoleResult

type compareRoleResult struct {
	RoleName       string
	UserIDsResult  compareIDsResult
	GroupIDsResult compareIDsResult
}

type compareIDsResult struct {
	CreatedIDs  []string
	DeletedIDs  []string
	RetainedIDs []string
}

func (c compareResult) hasContent() bool {
	for _, roleResult := range c {
		if roleResult.hasContent() {
			return true
		}
	}
	return false
}

func (c compareRoleResult) hasContent() bool {
	return c.UserIDsResult.hasContent() || c.GroupIDsResult.hasContent()
}

func (c compareIDsResult) hasContent() bool {
	return len(c.CreatedIDs) > 0 || len(c.DeletedIDs) > 0 || len(c.RetainedIDs) > 0
}

func (c compareResult) show() {
	for _, roleResult := range c {
		if roleResult.hasContent() {
			fmt.Println("role: " + colorOfStrong.SprintFunc()(roleResult.RoleName))
			roleResult.show()
		}
	}
}

func (c compareRoleResult) show() {
	printIndent(2)
	fmt.Println("users:")
	if c.UserIDsResult.hasContent() {
		c.UserIDsResult.show()
	} else {
		printIndent(4)
		fmt.Println("(none)")
	}
	fmt.Println()
	printIndent(2)
	fmt.Println("groups:")
	if c.GroupIDsResult.hasContent() {
		c.GroupIDsResult.show()
	} else {
		printIndent(4)
		fmt.Println("(none)")
	}
	fmt.Println()
}

func (c compareIDsResult) show() {
	if len(c.RetainedIDs) > 0 {
		showAsCreated(4, c.RetainedIDs...)
	}
	if len(c.CreatedIDs) > 0 {
		showAsCreated(4, c.CreatedIDs...)
	}
	if len(c.DeletedIDs) > 0 {
		showAsDeleted(4, c.DeletedIDs...)
	}
}

func Compare(oldTeam, newTeam models.Team) compareResult {
	roleResults := compareResult{}

	for roleName, oldRule := range oldTeam.Auth {
		var roleResult compareRoleResult

		newRule, exists := newTeam.Auth[roleName]
		if exists {
			roleResult = compareRule(oldRule, newRule)
		} else {
			roleResult = compareRoleResult{
				RoleName:       roleName,
				UserIDsResult:  compareIDsResult{DeletedIDs: oldRule.Users},
				GroupIDsResult: compareIDsResult{DeletedIDs: oldRule.Groups},
			}
		}
		roleResults = append(roleResults, roleResult)
	}

	for roleName, newRule := range newTeam.Auth {
		_, exists := oldTeam.Auth[roleName]
		if !exists {
			roleResult := compareRoleResult{
				RoleName:       roleName,
				UserIDsResult:  compareIDsResult{CreatedIDs: newRule.Users},
				GroupIDsResult: compareIDsResult{CreatedIDs: newRule.Groups},
			}
			roleResults = append(roleResults, roleResult)
		}
	}

	roleResults.show()

	return roleResults
}

func compareRule(oldRule, newRule *models.AuthRule) compareRoleResult {
	return compareRoleResult{
		RoleName:       oldRule.RoleName,
		UserIDsResult:  compareIds(oldRule.Users, newRule.Users),
		GroupIDsResult: compareIds(oldRule.Groups, newRule.Groups),
	}
}

func compareIds(oldIds, newIds []string) compareIDsResult {
	oldIdsSet, newIdsSet := stringSet.New(oldIds...), stringSet.New(newIds...)
	deletedIdsSet, createdIdsSet, retainedIdsSet := stringSet.Partition(oldIdsSet, newIdsSet)
	return compareIDsResult{createdIdsSet.Array(), deletedIdsSet.Array(), retainedIdsSet.Array()}
}

func showAsCreated(indentLevel int, values ...string) {
	showWithColor(indentLevel, "+", colorOfCreated, values...)
}

func showAsDeleted(indentLevel int, values ...string) {
	showWithColor(indentLevel, "-", colorOfDeleted, values...)
}

func showAsRetained(indentLevel int, values ...string) {
	showWithColor(indentLevel, " ", colorOfRetained, values...)
}

func showWithColor(indentLevel int, prefix string, color *color.Color, values ...string) {
	for _, value := range values {
		printIndent(indentLevel)
		color.Printf("%s %s\n", prefix, value)
	}
}

func printIndent(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}
