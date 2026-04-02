package utils

import "encoding/json"

func BuildPrompt(schema map[string]any) string {
	schemaBytes, _ := json.MarshalIndent(schema, "", "  ")

	return `You are a JSON generator.

Task:
Generate ONLY a valid JSON object from the schema.

STRICT RULES:
- DO NOT include "type", "properties", or schema keywords
- DO NOT wrap values inside objects unless required
- Arrays must contain direct values (e.g., ["a","b"])
- Objects must contain only actual fields, not schema info
- Output ONLY JSON

Example:
Schema:
{"type":"object","properties":{"tags":{"type":"array","items":{"type":"string"}}}}

Output:
{"tags":["tag1","tag2"]}

Now generate for this schema:

` + string(schemaBytes)
}
