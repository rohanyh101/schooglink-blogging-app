package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/schooglink_assignment/helpers"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use the Authorization header instead of token
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header found"})
			c.Abort()
			return
		}

		// Remove "Bearer " from the token string
		token := clientToken[len("Bearer "):]

		claims, err := helpers.ValidateUserToken(token)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		// log.Printf("%+v", claims)

		c.Next()
	}
}
