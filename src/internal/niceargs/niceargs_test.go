package niceargs

import (
	"testing"
)

func TestList_NoneGiven(t *testing.T) {
	result := List([]string{})

	if !testEq(result, []string{}) {
		t.Errorf("If given no arguments, this function should return no arguments")
	}
}

func TestList_SelfContainedFlagGiven(t *testing.T) {
	result := List([]string{"--config=some_file.json"})

	if !testEq(result, []string{"--config=some_file.json"}) {
		t.Errorf("If given a self-contained flag, it should return it.")
	}
}

func TestList_TwoArgFlagGiven(t *testing.T) {
	result := List([]string{"--config", "some_file.json"})

	if !testEq(result, []string{"--config", "some_file.json"}) {
		t.Errorf(`If given a two argument flag, it should return the args in the
             same order.`)
	}
}

func TestList_ArgsWithoutFlag(t *testing.T) {
	result := List([]string{"some-argument-1", "some-argument-2"})

	if !testEq(result, []string{"some-argument-1", "some-argument-2"}) {
		t.Errorf(`If given arguments that are not flags, it should return the args
             in the same order.`)
	}
}

func TestList_ArgsAfterFlags(t *testing.T) {
	arguments := []string{"--config=some_file.json",
		"--some-flag",
		"some-value",
		"some-argument"}
	result := List(arguments)

	if !testEq(result, arguments) {
		t.Errorf(`If given flags before arguments, it should return everything in
             the same order.`)
	}
}

func TestList_ArgsBeforeAndAfterFlags(t *testing.T) {
	arguments := []string{"some-argument-1",
		"--config=some_file.json",
		"--some-flag",
		"some-value",
		"some-argument-2"}
	expected := []string{"--config=some_file.json",
		"--some-flag",
		"some-value",
		"some-argument-1",
		"some-argument-2"}
	result := List(arguments)

	if !testEq(result, expected) {
		t.Errorf(`If given arguments flags, the function should reorder them so the
             flags are first`)
	}
}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
