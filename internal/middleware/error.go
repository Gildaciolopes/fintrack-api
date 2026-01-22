package middleware

import (
	"net/http"

	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					Success: false,
					Error:   "Internal server error",
					Message: "An unexpected error occurred",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
