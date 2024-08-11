package helpers

import (
	"net/http"

	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/middleware"
	"github.com/gin-gonic/gin"
)

// GetAuthUser returns the authenticated user details from the Gin context
func GetAuthUser(c *gin.Context) *middleware.AuthUser {
	authUser, exists := c.Get("authUser")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get the user",
		})
		return nil
	}

	if user, ok := authUser.(middleware.AuthUser); ok {
		return &user
	}

	return nil
}
