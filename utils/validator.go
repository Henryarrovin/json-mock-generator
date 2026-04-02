package utils

import "encoding/json"

func IsValidJSON(input string) bool {
	var js any
	return json.Unmarshal([]byte(input), &js) == nil
}
