package projections

import (
	"fmt"
	"testing"
)

func TestFind_NotPresent(t *testing.T) {
	result := Find(nil, "some-key")

	if result != nil {
		t.Errorf("Function should return nothing if no projection config.")
	}
}

func TestFind_PresentSubkeys_WithoutAlternatesSubkey(t *testing.T) {
	projections := map[string]interface{}{
		"app/jobs/*.rb": map[string]interface{}{
			"not-alternate": "dont-care-please-stop",
		},
	}
	result := Find(projections, "alternate")

	if len(result) != 0 {
		t.Errorf("Function should return 1 result but returned %d", len(result))
	}
}

func TestFind_PresentSubkeys_WithAlternatesSubkey(t *testing.T) {
	projections := map[string]interface{}{
		"app/models/*.rb": map[string]interface{}{
			"alternate": "spec/models/{}_spec.rb",
		},
		"app/jobs/*.rb": map[string]interface{}{
			"not-alternate": "dont-care-please-stop",
		},
	}
	result := Find(projections, "alternate")

	if len(result) != 1 {
		t.Errorf("Function should return 1 result but returned %d", len(result))
	}

	if result["app/models/*.rb"] == nil {
		t.Errorf(`Function should return a map that includes the key
             'app/models/*.rb'.`)
	}

	if result["app/jobs/*.rb"] != nil {
		t.Errorf(`Function should only return a map that includes the key
             'app/models/*.rb'.`)
	}
}

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
