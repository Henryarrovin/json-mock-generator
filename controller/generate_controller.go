package controller

import (
	"encoding/json"
	"json-mock-generator/client"
	"json-mock-generator/utils"
	"json-mock-generator/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GenerateData(c *gin.Context, schema map[string]any) (any, error) {
	corrID := c.GetString("CorrelationID")
	prompt := utils.BuildPrompt(schema)

	for i := 0; i < 3; i++ {
		result, err := client.CallOllama(prompt)
		if err != nil {
			logger.Log.Sugar().Warn("warn.generate_controller.call_ollama.tried_time: ",
				i,
				zap.String("correlation_id", corrID),
			)
			continue
		}

		cleanData := utils.ExtractJSON(result)
		if cleanData == "" {
			logger.Log.Sugar().Warn("warn.generate_controller.extract_json.tried_time: ",
				i,
				zap.String("correlation_id", corrID),
			)
			continue
		}

		if utils.IsSchemaOutput(cleanData) {
			logger.Log.Sugar().Warn("warn.generate_controller.is_schema_validation.tried_time: ",
				i,
				zap.String("correlation_id", corrID),
			)
			continue
		}

		if utils.IsValidJSON(cleanData) {
			var parsed any
			json.Unmarshal([]byte(cleanData), &parsed)
			logger.Log.Sugar().Info("info.generate_controller.generate_success",
				zap.String("correlation_id", corrID),
				zap.String("clean_data", cleanData),
			)
			return parsed, nil
		}
	}

	logger.Log.Sugar().Info("info.generate_controller.fallback_data",
		zap.String("correlation_id", corrID),
	)

	return utils.GenerateFallback(schema), nil
}
