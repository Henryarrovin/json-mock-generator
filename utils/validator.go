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

func ExtractJSON(input string) string {
	start := strings.Index(input, "{")
	end := strings.LastIndex(input, "}")

	if start == -1 || end == -1 || start >= end {
		return ""
	}

	return input[start : end+1]
}
