package utils

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateFallback(schema map[string]interface{}) interface{} {
	t, ok := schema["type"].(string)
	if !ok {
		return nil
	}

	switch t {

	case "string":
		return "sample_text"

	case "integer":
		return rng.Intn(100) + 1

	case "boolean":
		return rng.Intn(2) == 1

	case "array":
		items, ok := schema["items"].(map[string]interface{})
		if !ok {
			return []interface{}{}
		}

		return []interface{}{
			GenerateFallback(items),
			GenerateFallback(items),
		}

	case "object":
		props, ok := schema["properties"].(map[string]interface{})
		if !ok {
			return map[string]interface{}{}
		}

		result := make(map[string]interface{})

		for key, val := range props {
			if subSchema, ok := val.(map[string]interface{}); ok {
				result[key] = GenerateFallback(subSchema)
			}
		}

		return result
	}

	return nil
}
