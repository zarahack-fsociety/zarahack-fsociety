package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trephy/unit/internal/middleware"
)

func NewRouter(h *HealthHandler) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	r.GET("/health", h.Health)
	r.GET("/ready", h.Ready)

	return r
}
