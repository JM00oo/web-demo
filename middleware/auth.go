package middleware

import (
	// "database/sql"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/web-demo/model"
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
		userStore := model.NewUserStore()
		user, err := userStore.GetByToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errorMsg": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("userName", user.Username)
		c.Next()
	}
}
