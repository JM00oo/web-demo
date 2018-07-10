package middleware

import (
	// "database/sql"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequired : check tenant's session token in db
func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"errorMsg": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
