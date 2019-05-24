package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Team struct {
	ID   int                 `json:"id"`
	Name string              `json:"name"`
	Auth map[string]AuthRule `json:"auth"`
}

type AuthRule struct {
	Users  []string `json:"users"`
	Groups []string `json:"groups"`
}

func loadTextFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "args: <existing config (json)> <new config (yaml)>")
		os.Exit(1)
	}

	jsonData, err := loadTextFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	yamlData, err := loadTextFromFile(os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println(string(yamlData))

	team := Team{}
	json.Unmarshal(jsonData, &team)
	fmt.Printf("%#v\n", team)
}
