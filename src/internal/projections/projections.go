package projections

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func AlternateOf(projections map[string]interface{},
	file string) (string, error) {
	if projections == nil {
		return "", errors.New("No projections provided.")
	}

	candidates := keyCandidates(projections, file)
	relevantProjections := filterProjections(projections, candidates)
	results := replaceTemplate(relevantProjections, candidates, file, "alternate")

	if len(results) == 0 {
		return "", errors.New(fmt.Sprintf("No alternate found for '%s'.", file))
	}

	return strings.Join(results, " "), nil
}

func Find(projections map[string]interface{}, subkey string) []string {
	if projections == nil {
		return nil
	}

	results := make([]string, 0, len(projections))

	for _, object := range projections {
		o := object.(map[string]interface{})[subkey]

		if o != nil {
			results = append(results, fmt.Sprintf("%v", o))
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

func filterProjections(projections map[string]interface{},
	wantedKeys []string) map[string]interface{} {
	results := map[string]interface{}{}

	for key, object := range projections {
		for _, candidate := range wantedKeys {
			if key == candidate {
				results[key] = object
				break
			}
		}
	}

	return results
}

func keyCandidates(projections map[string]interface{}, file string) []string {
	candidates := make([]string, 0, len(projections))

	for key, _ := range projections {
		regexpReadyKey := strings.Replace(key, "*", ".*", -1)
		matched, _ := regexp.Match(regexpReadyKey, []byte(file))

		if matched {
			candidates = append(candidates, key)
		}
	}

	return candidates
}

func replaceTemplate(projections map[string]interface{},
	candidates []string,
	file string,
	subkey string) []string {
	results := make([]string, 0, len(projections))

	for _, c := range candidates {
		candidate := strings.Replace(c, "*", "(?P<Template>.*)", -1)
		re := regexp.MustCompile(candidate)
		match := re.FindStringSubmatch(file)

		for _, p := range Find(projections, subkey) {
			results = append(results,
				strings.Replace(p, "{}", match[re.SubexpIndex("Template")], -1))
		}
	}

	return results
}
