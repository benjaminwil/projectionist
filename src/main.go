package main

import (
	"encoding/json"
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
	projections, err := Read(ioutil.ReadFile, file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	form := make(map[string]string)

	for key, value := range projections {
		form[key] = value.(string)
	}

	return fmt.Sprintf("%v", form)
}

func Read(f func(string) ([]byte, error), file *string) (map[string]interface{}, error) {
	if contents, err := f(*file); err == nil {

		var data map[string]interface{}
		e := json.Unmarshal(contents, &data)

		if e == nil {
			return data, nil
		} else {
			return nil, e
		}
	} else {
		return nil, err
	}
}
