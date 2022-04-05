package projections

import (
	"encoding/json"
)

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
