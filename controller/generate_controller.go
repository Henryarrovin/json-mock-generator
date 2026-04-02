package controller

import (
	"encoding/json"
	"errors"
	"json-mock-generator/client"
	"json-mock-generator/utils"
)

func GenerateData(schema map[string]any) (any, error) {
	prompt := utils.BuildPrompt(schema)

	for i := 0; i < 3; i++ {
		result, err := client.CallOllama(prompt)
		if err != nil {
			continue
		}

		if utils.IsValidJSON(result) {
			var parsed any
			json.Unmarshal([]byte(result), &parsed)
			return parsed, nil
		}
	}

	return nil, errors.New("failed to generate valid JSON")
}
