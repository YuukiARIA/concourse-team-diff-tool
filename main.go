package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/YuukiARIA/concourse-team-diff-tool/models"
)

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

	team := models.NewFromJSON(jsonData)
	fmt.Printf("%#v\n", team)

	newTeam := LoadYAML(yamlData)
	fmt.Printf("%#v\n", newTeam)
}
