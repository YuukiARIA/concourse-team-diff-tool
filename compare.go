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
			roleResult = models.CompareRoleResult{
				RoleName:       roleName,
				Deleted:        true,
				UserIDsResult:  newCompareIDsResult(nil, oldRule.Users, nil),
				GroupIDsResult: newCompareIDsResult(nil, oldRule.Groups, nil),
			}
		}
		roleResults = append(roleResults, roleResult)
	}

	for roleName, newRule := range newTeam.Auth {
		_, exists := oldTeam.Auth[roleName]
		if !exists {
			roleResult := models.CompareRoleResult{
				RoleName:       roleName,
				Created:        true,
				UserIDsResult:  newCompareIDsResult(newRule.Users, nil, nil),
				GroupIDsResult: newCompareIDsResult(newRule.Groups, nil, nil),
			}
			roleResults = append(roleResults, roleResult)
		}
	}

	return models.CompareResult{TeamName: oldTeam.Name, Results: roleResults}
}

func compareRule(oldRule, newRule *models.AuthRule) models.CompareRoleResult {
	return models.CompareRoleResult{
		RoleName:       oldRule.RoleName,
		UserIDsResult:  compareIds(oldRule.Users, newRule.Users),
		GroupIDsResult: compareIds(oldRule.Groups, newRule.Groups),
	}
}

func compareIds(oldIds, newIds []string) models.CompareIDsResult {
	oldIdsSet, newIdsSet := stringset.New(oldIds...), stringset.New(newIds...)
	deletedIdsSet, createdIdsSet, retainedIdsSet := stringset.Partition(oldIdsSet, newIdsSet)
	return newCompareIDsResult(createdIdsSet.Array(), deletedIdsSet.Array(), retainedIdsSet.Array())
}

func newCompareIDsResult(createdIDs, deletedIDs, retainedIDs []string) models.CompareIDsResult {
	empty := make([]string, 0)
	if createdIDs == nil {
		createdIDs = empty
	}
	if deletedIDs == nil {
		deletedIDs = empty
	}
	if retainedIDs == nil {
		retainedIDs = empty
	}
	return models.CompareIDsResult{CreatedIDs: createdIDs, DeletedIDs: deletedIDs, RetainedIDs: retainedIDs}
}
