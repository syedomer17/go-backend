package middleware

import (
	"net/http"
	"task-manager/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		token, err := c.Cookie("accessToken")

		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		clames, err := utils.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"error": "Invalid Token",
			})
			c.Abort()
			return
		}

		c.Set("userId",clames.UserID)
		
		c.Next()
	}
}