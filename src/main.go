package main

import (
	"flag"
	"fmt"
	"internal/projections"
	"io/ioutil"
	"os"
	"strings"
)

var (
	alternateCommand = flag.NewFlagSet("alternate", flag.ExitOnError)
	config           *string
	subcommands      = map[string]*flag.FlagSet{
		alternateCommand.Name(): alternateCommand,
	}
)

func main() {
	configureFlags()
	exitIfInvalidArguments()

	command := subcommands[os.Args[1]]
	command.Parse(os.Args[2:])

	fmt.Println(read(config))
	os.Exit(0)
}

func exitIfInvalidArguments() {
	if (len(os.Args) < 2) || subcommands[os.Args[1]] == nil {
		fmt.Println("Valid subcommands:", subcommandList())
		os.Exit(1)
	}
}

func subcommandList() string {
	list := make([]string, 0, len(subcommands))

	for name := range subcommands {
		list = append(list, name)
	}

	return strings.Join(list, ", ")
}

func configureFlags() {
	for _, flagSet := range subcommands {
		config = flagSet.String(
			"config",
			fmt.Sprintf("%s/.projections.json", pwd()),
			"Provide a JSON file to read projections from.",
		)
	}
}

func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return path
}

func read(config *string) string {
	projections := projections.Read(ioutil.ReadFile, config)

	if projections == nil {
		fmt.Println("No projections in:", *config)
		os.Exit(1)
	}

	form := make(map[string]string)

	for key, value := range projections {
		form[key] = value.(string)
	}

	return fmt.Sprintf("%v", form)
}
