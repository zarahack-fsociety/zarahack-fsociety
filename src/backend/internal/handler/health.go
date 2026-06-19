package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
)

type HealthHandler struct {
	DB  *sql.DB
	RDB *goredis.Client
}

func NewHealthHandler(db *sql.DB, rdb *goredis.Client) *HealthHandler {
	return &HealthHandler{DB: db, RDB: rdb}
}

type healthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Redis    string `json:"redis"`
}

func (h *HealthHandler) Health(c *gin.Context) {
	resp := healthResponse{Status: "ok", Database: "ok", Redis: "ok"}
	statusCode := http.StatusOK

	if err := h.DB.PingContext(c.Request.Context()); err != nil {
		resp.Database = "unavailable"
		statusCode = http.StatusServiceUnavailable
		slog.Warn("database health check failed", "error", err)
	}

	if err := h.RDB.Ping(c.Request.Context()).Err(); err != nil {
		resp.Redis = "unavailable"
		statusCode = http.StatusServiceUnavailable
		slog.Warn("redis health check failed", "error", err)
	}

	if resp.Database != "ok" || resp.Redis != "ok" {
		resp.Status = "degraded"
	}

	c.JSON(statusCode, resp)
}

func (h *HealthHandler) Ready(c *gin.Context) {
	resp := healthResponse{Status: "ok", Database: "ok", Redis: "ok"}

	if err := h.DB.PingContext(c.Request.Context()); err != nil {
		resp.Database = "unavailable"
		resp.Status = "unavailable"
		slog.Warn("readiness check: database failed", "error", err)
		c.JSON(http.StatusServiceUnavailable, resp)
		return
	}

	if err := h.RDB.Ping(c.Request.Context()).Err(); err != nil {
		resp.Redis = "unavailable"
		resp.Status = "unavailable"
		slog.Warn("readiness check: redis failed", "error", err)
		c.JSON(http.StatusServiceUnavailable, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
