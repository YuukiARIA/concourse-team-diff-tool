package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/YuukiARIA/glanceable/formatter"
	"github.com/YuukiARIA/glanceable/models"
	"github.com/jessevdk/go-flags"
)

type options struct {
	ShowVersion    func() `short:"v" long:"version" description:"show version"`
	ConfigFileName string `short:"c" long:"config" description:"team config file (yaml)" required:"yes"`
	Format         string `short:"f" long:"format" default:"default" choice:"default" choice:"json" choice:"yaml" choicedescription:"output format (default, json, yaml)"`
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

	jsonData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	yamlData, err := ioutil.ReadFile(opts.ConfigFileName)
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

	formatter.FormatResult(Compare(oldTeam, newTeam), opts.Format)
}
