package models

type CompareResult struct {
	TeamName string              `yaml:"team_name"`
	Results  []CompareRoleResult `yaml:"results"`
}

type CompareRoleResult struct {
	RoleName       string           `yaml:"role_name"`
	Created        bool             `yaml:"created"`
	Deleted        bool             `yaml:"deleted"`
	UserIDsResult  CompareIDsResult `yaml:"user_ids_result"`
	GroupIDsResult CompareIDsResult `yaml:"group_ids_result"`
}

type CompareIDsResult struct {
	CreatedIDs  []string `yaml:"created_ids"`
	DeletedIDs  []string `yaml:"deleted_ids"`
	RetainedIDs []string `yaml:"retained_ids"`
}

func (c CompareResult) HasContent() bool {
	for _, roleResult := range c.Results {
		if roleResult.HasContent() {
			return true
		}
	}
	return false
}

func (c CompareRoleResult) HasContent() bool {
	return c.UserIDsResult.HasContent() || c.GroupIDsResult.HasContent()
}

func (c CompareIDsResult) HasContent() bool {
	return len(c.CreatedIDs) > 0 || len(c.DeletedIDs) > 0 || len(c.RetainedIDs) > 0
}
