package userRoutes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(user *gin.RouterGroup) {
	user.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User profile"})
	})

	// Add more customer page routes here
}
