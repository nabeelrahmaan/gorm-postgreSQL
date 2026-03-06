package middlware

import (
	"authentication/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = utils.AccessSecret

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "authorization header missing",})
			c.Abort()
			return 
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 {
			c.JSON(401, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}

		tokenstring := parts[1]

		token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token",})
			c.Abort()
			return 
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}