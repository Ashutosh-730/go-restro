package adminRoutes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(admin *gin.RouterGroup) {
	admin.POST("/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin dashboard"})
	})

	// Add more admin panel routes here
}
