package formatter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/YuukiARIA/glanceable/models"
	"github.com/logrusorgru/aurora"
	yaml "gopkg.in/yaml.v2"
)

const (
	colorOfCreated      = aurora.GreenFg
	colorOfDeleted      = aurora.RedFg
	colorOfRetained     = aurora.WhiteFg
	colorOfCreatedRole  = aurora.GreenFg | aurora.BrightFg | aurora.BoldFm
	colorOfDeletedRole  = aurora.RedFg | aurora.BrightFg | aurora.BoldFm
	colorOfRetainedRole = aurora.WhiteFg | aurora.BrightFg | aurora.BoldFm
	colorOfTeam         = aurora.WhiteFg | aurora.BrightFg | aurora.BoldFm
)

const (
	formatDefault = "default"
	formatJSON    = "json"
	formatYAML    = "yaml"
)

func FormatResult(result models.CompareResult, format string) {
	switch strings.ToLower(format) {
	case formatDefault:
		showDefaultFormat(result)
	case formatJSON:
		showJSONFormat(result)
	case formatYAML:
		showYAMLFormat(result)
	}
}

func showDefaultFormat(result models.CompareResult) {
	showCompareResult(result)
}

func showJSONFormat(result models.CompareResult) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(&result)
}

func showYAMLFormat(result models.CompareResult) error {
	bytes, err := yaml.Marshal(&result)
	if err != nil {
		return err
	}
	fmt.Print(string(bytes))
	return nil
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

func showWithColor(indentLevel int, prefix string, color aurora.Color, values ...string) {
	for _, value := range values {
		printIndent(indentLevel)
		fmt.Println(aurora.Colorize(fmt.Sprintf("%s %s", prefix, value), color))
	}
}

func printIndent(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func showCompareResult(c models.CompareResult) {
	fmt.Println("team: " + aurora.Colorize(c.TeamName, colorOfTeam).String())
	fmt.Println()

	for _, roleResult := range c.Results {
		if roleResult.HasContent() {
			c := getRoleColor(roleResult.Created, roleResult.Deleted)
			fmt.Println("role: " + aurora.Colorize(roleResult.RoleName, c).String())
			showCompareRoleResult(roleResult)
		}
	}
}

func showCompareRoleResult(c models.CompareRoleResult) {
	printIndent(2)
	fmt.Println("users:")
	if c.UserIDsResult.HasContent() {
		showCompareIDsResult(c.UserIDsResult)
	} else {
		printIndent(4)
		fmt.Println("(none)")
	}
	fmt.Println()
	printIndent(2)
	fmt.Println("groups:")
	if c.GroupIDsResult.HasContent() {
		showCompareIDsResult(c.GroupIDsResult)
	} else {
		printIndent(4)
		fmt.Println("(none)")
	}
	fmt.Println()
}

func showCompareIDsResult(c models.CompareIDsResult) {
	if len(c.RetainedIDs) > 0 {
		showAsRetained(4, c.RetainedIDs...)
	}
	if len(c.CreatedIDs) > 0 {
		showAsCreated(4, c.CreatedIDs...)
	}
	if len(c.DeletedIDs) > 0 {
		showAsDeleted(4, c.DeletedIDs...)
	}
}

func getRoleColor(created, deleted bool) aurora.Color {
	if created {
		return colorOfCreatedRole
	} else if deleted {
		return colorOfDeletedRole
	}
	return colorOfRetainedRole
}
