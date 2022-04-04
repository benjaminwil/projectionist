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
	projections := Read(ioutil.ReadFile, file)

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

func Has(projections map[string]interface{}, subkey string) map[string]interface{} {
	if projections == nil {
		return nil
	}

	results := map[string]interface{}{}

	for key, object := range projections {
		if object.(map[string]interface{})[subkey] != nil {
			results[key] = object
		}
	}

	return results
}

func Read(f func(string) ([]byte, error), file *string) map[string]interface{} {
	contents, err := f(*file)

	if err != nil {
		return nil
	}

	var data map[string]interface{}
	if e := json.Unmarshal(contents, &data); e != nil {
		return nil
	}

	return data
}
