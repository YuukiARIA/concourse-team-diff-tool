package main

import (
	"fmt"

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

func LoadYAML(yamlData []byte) {
	teamDraft := teamDraft{}
	yaml.Unmarshal(yamlData, &teamDraft)
	fmt.Printf("TEAM DRAFT\n%#v\n", teamDraft)
}
