package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestRead_NoFilePresent(t *testing.T) {
	expectedError := fmt.Errorf("no such file")
	file := pointer("./.projections.json")

	f := func(filename string) ([]byte, error) {
		return make([]byte, 0, 48), expectedError
	}

	result, err := Read(f, file)

	if result != nil {
		t.Errorf("Function should return 'no file' but returned %s", result)
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Expected a 'no such file' error but didn't get one.")
	}
}

func TestRead_FilePresent(t *testing.T) {
	file := pointer("./.projections.json")
	fileContents := "{}"
	f := func(filename string) ([]byte, error) {
		return []byte(fileContents), nil
	}

	result, _ := Read(f, file)

	if result == nil {
		t.Errorf("function should return JSON interface for '%s' but didn't",
			fileContents)
	}
}

func pointer(s string) *string {
	var p *string
	p = &s
	return p
}
