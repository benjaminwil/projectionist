package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

type DummyReadFile struct {
	Str string
}

func (f DummyReadFile) ReadFile(filename string) ([]byte, error) {
	buffer := bytes.NewBufferString(f.Str)
	return ioutil.ReadAll(buffer)
}

func TestReadProjection_NoFilePresent(t *testing.T) {
	expectedError := fmt.Errorf("no such file")
	noFileToRead := func(filename string) ([]byte, error) {
		return make([]byte, 0, 48), expectedError
	}
	result, err := ReadProjections(noFileToRead)

	if result != "no file" {
		t.Errorf("Function should return 'no file' but returned %s", result)
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Expected a 'no such file' error but didn't get one.")
	}
}

func TestReadProjections_FilePresent(t *testing.T) {
	dummy_ioutil := DummyReadFile{Str: "{}"}
	result, _ := ReadProjections(dummy_ioutil.ReadFile)

	if result != "{}" {
		t.Errorf("function should return '{}' but returned %s", result)
	}
}
