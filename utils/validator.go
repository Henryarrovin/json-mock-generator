package utils

import (
	"encoding/json"
	"strings"
)

func IsValidJSON(input string) bool {
	var js any
	return json.Unmarshal([]byte(input), &js) == nil
}

func IsSchemaOutput(result string) bool {
	return strings.Contains(result, `"type"`) ||
		strings.Contains(result, `"properties"`)
}
