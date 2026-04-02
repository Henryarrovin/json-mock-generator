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

		cleanData := utils.ExtractJSON(result)
		if cleanData == "" {
			continue
		}

		if utils.IsSchemaOutput(cleanData) {
			continue
		}

		if utils.IsValidJSON(cleanData) {
			var parsed any
			json.Unmarshal([]byte(cleanData), &parsed)
			return parsed, nil
		}
	}

	return utils.GenerateFallback(schema), nil
}
