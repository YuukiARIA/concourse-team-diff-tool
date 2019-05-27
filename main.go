package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/YuukiARIA/glanceable/models"
)

type options struct {
	ShowVersion    func() `short:"v" long:"version" description:"show version"`
	ConfigFileName string `short:"c" long:"config" description:"team config file (yaml)" required:"yes"`
	Format         string `short:"f" long:"format" default:"default" choice:"default" choice:"json" choice:"yaml" choicedescription:"output format (default, json, yaml)"`
}

const (
	formatDefault = "default"
	formatJSON    = "json"
	formatYAML    = "yaml"
)

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
	switch strings.ToLower(format) {
	case formatDefault:
		ShowDefaultFormat(result)
	case formatJSON:
		ShowJSONFormat(result)
	case formatYAML:
		ShowYAMLFormat(result)
	}
}

func main() {
	var opts options

	opts.ShowVersion = func() {
		fmt.Printf("%s v%s\n", filepath.Base(os.Args[0]), version)
		os.Exit(0)
	}

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

	oldTeam, err := models.NewFromJSON(jsonData)
	if err != nil {
		panic(err)
	}
	newTeam, err := LoadYAML(yamlData)
	if err != nil {
		panic(err)
	}

	showResult(Compare(oldTeam, newTeam), opts.Format)
}
