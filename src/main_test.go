package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestReadProjection_NoFilePresent(t *testing.T) {
	expectedError := fmt.Errorf("no such file")
	path := pointer("./.projections.json")

	f := func(filename string) ([]byte, error) {
		return make([]byte, 0, 48), expectedError
	}

	result, err := ReadProjections(f, path)

	if result != "no file" {
		t.Errorf("Function should return 'no file' but returned %s", result)
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Expected a 'no such file' error but didn't get one.")
	}
}

func TestReadProjections_FilePresent(t *testing.T) {
	json := "{}"
	path := pointer("./.projections.json")

	f := func(filename string) ([]byte, error) {
		return []byte(json), nil
	}

	result, _ := ReadProjections(f, path)

	if result != "{}" {
		t.Errorf("function should return '{}' but returned %s", result)
	}
}

func pointer(s string) *string {
	var p *string
	p = &s
	return p
}
