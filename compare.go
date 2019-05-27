package main

import (
	"sort"

	"github.com/YuukiARIA/glanceable/models"
	"github.com/YuukiARIA/glanceable/stringset"
)

func Compare(oldTeam, newTeam *models.Team) models.CompareResult {
	roleResults := make([]models.CompareRoleResult, 0)

	roleNameSet := stringset.New()
	for rn := range oldTeam.Auth {
		roleNameSet.Add(rn)
	}
	for rn := range newTeam.Auth {
		roleNameSet.Add(rn)
	}
	roleNames := roleNameSet.Array()
	sort.Strings(roleNames)

	for _, roleName := range roleNames {
		oldRule, oldExists := oldTeam.Auth[roleName]
		newRule, newExists := newTeam.Auth[roleName]

		var roleResult models.CompareRoleResult

		switch {
		case !oldExists: // means the role is newly defined
			roleResult = models.NewCompareRoleResultCreated(
				roleName,
				models.NewCompareIDsResult(newRule.Users, nil, nil),
				models.NewCompareIDsResult(newRule.Groups, nil, nil),
			)
		case !newExists: // means the role definition was deleted
			roleResult = models.NewCompareRoleResultDeleted(
				roleName,
				models.NewCompareIDsResult(nil, oldRule.Users, nil),
				models.NewCompareIDsResult(nil, oldRule.Groups, nil),
			)
		default:
			roleResult = compareRule(oldRule, newRule)
		}

		roleResults = append(roleResults, roleResult)
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
