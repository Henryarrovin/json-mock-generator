package utils

import "encoding/json"

func BuildPrompt(schema map[string]any) string {
	schemaBytes, _ := json.MarshalIndent(schema, "", "  ")

	return `You are a JSON generator.

Task:
Generate a sample JSON object that matches the given JSON schema.

STRICT RULES:
- Output ONLY valid JSON
- DO NOT return the schema
- DO NOT include "type", "properties", or schema keywords
- Generate actual values (names, numbers, booleans, arrays)
- Follow schema structure exactly

Example:
Schema:
{"type":"object","properties":{"name":{"type":"string"}}}

Output:
{"name":"John"}

Now generate for this schema:
` + string(schemaBytes)
}
