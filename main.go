package main

import (
	"json-mock-generator/service"
	"json-mock-generator/utils/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting JSON Mock Generator API ...")

	r := gin.Default()

	r.Use(logger.RequestLoggerMiddleware)

	r.POST("/generate", service.GenerateHandler)

	r.Run(":8080")
}
