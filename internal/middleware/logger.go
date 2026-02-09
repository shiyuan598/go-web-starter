package middleware

import (
	"time"

	"go-web-starter/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		cost := time.Since(start)

		logger.Log.Info("http request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("cost", cost),
			zap.String("ip", c.ClientIP()),
		)
	}
}
