package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeaders adds security headers to all responses
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) { 
		c.Header("X-Content-Type-Options", "nosniff")
		 
		c.Header("X-Frame-Options", "DENY")
		 
		c.Header("X-XSS-Protection", "1; mode=block")
		 
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		 
		c.Header("X-Powered-By", "")
		 
		c.Header("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
		 
		c.Header("X-DNS-Prefetch-Control", "off")
		
		c.Next()
	}
}

// RateLimitInfo adds rate limit information headers (for future implementation)
func RateLimitInfo() gin.HandlerFunc {
	return func(c *gin.Context) { 
		
		c.Next()
	}
}
