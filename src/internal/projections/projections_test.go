package projections

import (
	"fmt"
	"testing"
)

func TestAlternateOf_ProjectionsNotPresent(t *testing.T) {
	_, err := AlternateOf(nil, "path/to/file.txt")

	if err == nil {
		t.Errorf("Function should return an error if no projections given.")
	}
}

func TestAlternateOf_AlternateFound(t *testing.T) {
	projections := map[string]interface{}{
		"app/jobs/*.rb": map[string]interface{}{
			"alternate": "spec/jobs/{}_spec.rb",
		},
	}

	result, err := AlternateOf(projections, "app/jobs/my_file.rb")

	if result != "spec/jobs/my_file_spec.rb" {
		t.Errorf("Function should return 'spec/jobs/my_file_spec.rb' but returned %s", result)
	}

	if err != nil {
		t.Errorf("Function should not return an error but returned '%v'.", err)
	}
}

func TestAlternateOf_NoAlternateFound(t *testing.T) {
	projections := map[string]interface{}{
		"app/jobs/*.rb": map[string]interface{}{
			"alternate": "spec/jobs/{}_spec.rb",
		},
	}

	result, err := AlternateOf(projections, "app/models/my_file.rb")

	if result != "" {
		t.Errorf("Function should return an empty string but returns %s", result)
	}

	if err == nil {
		t.Errorf("Function should return an error but did not.")
	}
}

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

	if !testEq(result, []string{"spec/models/{}_spec.rb"}) {
		t.Errorf("Function should return result 'spec/models/{}_spec.rb' but returned %v", result)
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
