package main

import (
	"github.com/YuukiARIA/glanceable/models"
	"github.com/YuukiARIA/glanceable/stringset"
)

func Compare(oldTeam, newTeam *models.Team) models.CompareResult {
	roleResults := make([]models.CompareRoleResult, 0)

	for roleName, oldRule := range oldTeam.Auth {
		var roleResult models.CompareRoleResult

		newRule, exists := newTeam.Auth[roleName]
		if exists {
			roleResult = compareRule(oldRule, newRule)
		} else {
			roleResult = models.NewCompareRoleResultDeleted(
				roleName,
				models.NewCompareIDsResult(nil, oldRule.Users, nil),
				models.NewCompareIDsResult(nil, oldRule.Groups, nil),
			)
		}
		roleResults = append(roleResults, roleResult)
	}

	for roleName, newRule := range newTeam.Auth {
		_, exists := oldTeam.Auth[roleName]
		if !exists {
			roleResult := models.NewCompareRoleResultCreated(
				roleName,
				models.NewCompareIDsResult(newRule.Users, nil, nil),
				models.NewCompareIDsResult(newRule.Groups, nil, nil),
			)
			roleResults = append(roleResults, roleResult)
		}
	}

	return models.NewCompareResult(oldTeam.Name, roleResults)
}

func compareRule(oldRule, newRule *models.AuthRule) models.CompareRoleResult {
	return models.NewCompareRoleResultRetained(
		oldRule.RoleName,
		compareIds(oldRule.Users, newRule.Users),
		compareIds(oldRule.Groups, newRule.Groups),
	)
}

func compareIds(oldIds, newIds []string) models.CompareIDsResult {
	oldIdsSet, newIdsSet := stringset.New(oldIds...), stringset.New(newIds...)
	deletedIdsSet, createdIdsSet, retainedIdsSet := stringset.Partition(oldIdsSet, newIdsSet)
	return models.NewCompareIDsResult(createdIdsSet.Array(), deletedIdsSet.Array(), retainedIdsSet.Array())
}
