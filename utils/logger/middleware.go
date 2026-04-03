package logger

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLoggerMiddleware(c *gin.Context) {
	start := time.Now()

	correlationID := uuid.New().String()
	c.Set("CorrelationID", correlationID)

	c.Writer.Header().Set("X-Correlation-ID", correlationID)

	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // restore body

	c.Next()

	statusCode := c.Writer.Status()
	latency := time.Since(start)

	Log.Info("HTTP Request",
		zap.String("correlation_id", correlationID),
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", statusCode),
		zap.Duration("latency", latency),
		zap.ByteString("request_body", bodyBytes),
	)
}
