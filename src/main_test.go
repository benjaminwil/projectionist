package main

import (
	"fmt"
	"testing"
)

func TestRead_NoFilePresent(t *testing.T) {
	expectedError := fmt.Errorf("no such file")
	file := pointer("./.projections.json")

	f := func(filename string) ([]byte, error) {
		return make([]byte, 0, 48), expectedError
	}

	result := Read(f, file)

	if result != nil {
		t.Errorf("Function should return nothing but returned %s", result)
	}
}

func TestRead_FilePresent(t *testing.T) {
	file := pointer("./.projections.json")
	fileContents := "{}"
	f := func(filename string) ([]byte, error) {
		return []byte(fileContents), nil
	}

	result := Read(f, file)

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
