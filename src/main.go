package main

import (
	"flag"
	"fmt"
	"internal/niceargs"
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

	arguments := niceargs.List(os.Args[2:])

	command := subcommands[os.Args[1]]
	command.Parse(arguments)
	data := readProjections()

	exitIfNoFilesGiven(command)

	printAlternates(data, command)

	fmt.Println(`Something went wrong.
              If you received this, you may want to submit a bug report.`)
	os.Exit(1)
}

func exitIfInvalidArguments() {
	if (len(os.Args) < 2) || subcommands[os.Args[1]] == nil {
		fmt.Println("Valid subcommands:", subcommandList())
		os.Exit(1)
	}
}

func exitIfNoFilesGiven(command *flag.FlagSet) {
	if command.NArg() == 0 {
		fmt.Println("No files given.")
		os.Exit(1)
	}
}

func printAlternates(data map[string]interface{}, command *flag.FlagSet) {
	if os.Args[1] == "alternate" {
		results, err := projections.AlternateOf(data, command.Args()[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(results)
		os.Exit(0)
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

func readProjections() map[string]interface{} {
	projections := projections.Read(ioutil.ReadFile, config)

	if projections == nil {
		fmt.Println("No projections in:", *config)
		os.Exit(1)
	}

	return projections
}
