package utils

import "encoding/json"

func BuildPrompt(schema map[string]any) string {
	schemaBytes, _ := json.MarshalIndent(schema, "", "  ")

	return `You are a JSON generator.

	Your job is to CREATE DATA, not explain schemas.

	STRICT RULES:
	- Return ONLY valid JSON
	- DO NOT return schema
	- DO NOT include "type", "properties", or schema keywords
	- Generate actual values only

	TYPE MAPPING:
	- string → "example text"
	- integer → 25
	- boolean → true
	- array → ["item1","item2"]
	- object → nested JSON

	BAD OUTPUT (DO NOT DO THIS):
	{"type":"object","properties":{...}}

	GOOD OUTPUT:
	{"name":"John","age":25}

	Now generate data for this schema:

	` + string(schemaBytes)
}
