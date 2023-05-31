package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"strings"
)


func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the request contains a valid authentication token
		if !isValidAuthToken(c.Request) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the token is valid, call the next handler
		c.Next()
	}
}

func orderHandler(c *gin.Context) {
	// Handle the order request
}

func isValidAuthToken(r *http.Request) bool {
	token , err := extractToken(r)
	if err != nil || token != "mLGV?uxpQ0rfsy-))CGW.1!=r#dKOiZe5#G]|B4Xi@tPuVF2SdubREaTAK#}DO3I" {
		return false
	}
	return true
}

func extractToken(r *http.Request) (string, error) {
    // Get the value of the Authorization header
    authHeader := r.Header.Get("Authorization")

    // Check if the header is not present
    if authHeader == "" {
        return "", errors.New("Authorization header is not present")
    }

    // Parse the token from the header value
    // The header value typically starts with "Bearer " followed by the token string
    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
        return "", errors.New("Invalid Authorization header format")
    }

    return parts[1], nil
}