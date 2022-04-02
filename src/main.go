package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	projections, err := ReadProjections(ioutil.ReadFile)

	if err == nil {
		fmt.Println(projections)
		os.Exit(0)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadProjections(readfile func(string) ([]byte, error)) (string, error) {
	if contents, err := readfile("./.projections.json"); err == nil {
		return string(contents), nil
	} else {
		return "no file", err
	}
}
