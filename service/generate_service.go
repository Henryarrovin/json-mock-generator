package service

import (
	"json-mock-generator/controller"
	"json-mock-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateHandler(c *gin.Context) {
	var req models.GenerateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	data, err := controller.GenerateData(req.Schema)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, models.GenerateResponse{
	// 	Data: data,
	// })
	c.JSON(http.StatusOK, data)
}
