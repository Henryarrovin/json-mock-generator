package main

import (
	"json-mock-generator/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/generate", service.GenerateHandler)

	r.Run(":8080")
}
