package controller

import (
	"encoding/json"
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

		if utils.IsSchemaOutput(result) {
			continue
		}

		if utils.IsValidJSON(result) {
			var parsed any
			json.Unmarshal([]byte(result), &parsed)
			return parsed, nil
		}
	}

	return utils.GenerateFallback(schema), nil
}
