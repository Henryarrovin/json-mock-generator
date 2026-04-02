package utils

import "encoding/json"

func BuildPrompt(schema map[string]any) string {
	schemaBytes, _ := json.MarshalIndent(schema, "", "  ")

	return `You are a JSON generator.

Rules:
- Return ONLY valid JSON
- No explanation
- Follow schema strictly

Schema:
` + string(schemaBytes)
}
