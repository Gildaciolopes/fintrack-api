package handler

import (
	"net/http"
	"time"

	"github.com/Gildaciolopes/fintrack-backend/internal/models"
	"github.com/gin-gonic/gin"
)
 
type HealthHandler struct {
	version string
}
 
func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{version: version}
}
 
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthCheck{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
		Version:   h.version,
	})
}
