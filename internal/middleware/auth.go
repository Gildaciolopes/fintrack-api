package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Gildaciolopes/fintrack-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: jwtSecret,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false,
				Error:   "Authorization header required",
				Message: "Please provide a valid authentication token",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false,
				Error:   "Invalid authorization format",
				Message: "Authorization header must be in the format: Bearer <token>",
			})
			c.Abort()
			return
		}

		token := parts[1]

		user, err := m.validateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false,
				Error:   "Invalid or expired token",
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", user.ID)
		c.Set("user_email", user.Email)
		c.Set("user", user)

		c.Next()
	}
}

func (m *AuthMiddleware) validateToken(token string) (*models.AuthUser, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	message := parts[0] + "." + parts[1]
	signature := parts[2]

	if !m.verifySignature(message, signature) {
		return nil, fmt.Errorf("invalid token signature")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode token payload")
	}

	var user models.AuthUser
	if err := json.Unmarshal(payload, &user); err != nil {
		return nil, fmt.Errorf("failed to parse token claims")
	}

	return &user, nil
}

func (m *AuthMiddleware) verifySignature(message, signature string) bool {
	mac := hmac.New(sha256.New, []byte(m.jwtSecret))
	mac.Write([]byte(message))
	expectedSignature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

func GetUserID(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, fmt.Errorf("user not authenticated")
	}

	id, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user ID format")
	}

	return id, nil
}

func GetUser(c *gin.Context) (*models.AuthUser, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("user not authenticated")
	}

	authUser, ok := user.(*models.AuthUser)
	if !ok {
		return nil, fmt.Errorf("invalid user format")
	}

	return authUser, nil
}
