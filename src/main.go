package main

import (
	"flag"
	"fmt"
	"internal/projections"
	"io/ioutil"
	"os"
)

func main() {
	config := flag.String("config",
		fmt.Sprintf("%s/.projections.json", pwd()),
		"pass an explicit path to the preferred projections JSON")
	flag.Parse()

	fmt.Println(read(config))

	os.Exit(0)
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
