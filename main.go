package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/YuukiARIA/glanceable/models"
)

type options struct {
	ConfigFileName string `short:"c" long:"config" description:"team config file (yaml)" required:"yes"`
	Format         string `short:"f" long:"format" default:"default" choice:"default" choice:"json" choice:"yaml" choicedescription:"output format (default, json, yaml)"`
}

func loadTextFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

func loadTextFromReader(reader io.Reader) ([]byte, error) {
	return ioutil.ReadAll(reader)
}

func showResult(result compareResult, format string) {
	switch format {
	case "default":
		ShowDefaultFormat(result)
	case "json":
		ShowJSONFormat(result)
	case "yaml":
		ShowYAMLFormat(result)
	}
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	jsonData, err := loadTextFromReader(os.Stdin)
	if err != nil {
		panic(err)
	}
	yamlData, err := loadTextFromFile(opts.ConfigFileName)
	if err != nil {
		panic(err)
	}

	oldTeam := models.NewFromJSON(jsonData)
	newTeam := LoadYAML(yamlData)

	showResult(Compare(oldTeam, newTeam), opts.Format)
}
