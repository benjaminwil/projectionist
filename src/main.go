package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := flag.String("path",
		fmt.Sprintf("%s/.projections.json", pwd()),
		"pass an explicit path to the preferred projections JSON")

	flag.Parse()

	projections, err := ReadProjections(ioutil.ReadFile, path)
	if err == nil {
		fmt.Println(projections)
		os.Exit(0)
	} else {
		fmt.Println(err)
		os.Exit(1)
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

func ReadProjections(readfile func(string) ([]byte, error),
	path *string) (string, error) {
	if contents, err := readfile(*path); err == nil {
		return string(contents), nil
	} else {
		return "no file", err
	}
}
