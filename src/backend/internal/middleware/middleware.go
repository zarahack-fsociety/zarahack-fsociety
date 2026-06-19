package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	start := time.Now()

	c.Next()

	slog.Info("http request",
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
		"status", c.Writer.Status(),
		"duration", time.Since(start).String(),
	)
}

func Recoverer(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			slog.Error("panic recovered", "error", rec, "path", c.Request.URL.Path)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "internal server error",
			})
		}
	}()
	c.Next()
}
