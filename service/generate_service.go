package service

import (
	"json-mock-generator/controller"
	"json-mock-generator/models"
	"json-mock-generator/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GenerateHandler(c *gin.Context) {
	corrID := c.GetString("CorrelationID")
	var req models.GenerateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Sugar().Error("err.generate_service.invalid_request_error",
			zap.String("correlation_id", corrID),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	data, err := controller.GenerateData(c, req.Schema)
	if err != nil {
		logger.Log.Sugar().Error("err.generate_service.get_data_error",
			zap.String("correlation_id", corrID),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.GenerateResponse{
		Data: data,
	})
}
