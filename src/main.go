package main

import (
	"flag"
	"fmt"
	"internal/projections"
	"io/ioutil"
	"os"
)

func main() {
	file := flag.String("file",
		fmt.Sprintf("%s/.projections.json", pwd()),
		"pass an explicit path to the preferred projections JSON")

	flag.Parse()

	fmt.Println(read(file))

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

func read(file *string) string {
	projections := projections.Read(ioutil.ReadFile, file)

	if projections == nil {
		fmt.Println("No projections in:", *file)
		os.Exit(1)
	}

	form := make(map[string]string)

	for key, value := range projections {
		form[key] = value.(string)
	}

	return fmt.Sprintf("%v", form)
}
