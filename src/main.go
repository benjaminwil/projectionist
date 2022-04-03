package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := flag.String("file",
		fmt.Sprintf("%s/.projections.json", pwd()),
		"pass an explicit path to the preferred projections JSON")

	flag.Parse()

	projections, err := Read(ioutil.ReadFile, file)
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

func Read(f func(string) ([]byte, error), file *string) (string, error) {
	if contents, err := f(*file); err == nil {
		return string(contents), nil
	} else {
		return "no file", err
	}
}
