package middlware

import "github.com/gin-gonic/gin"

func RoleMiddleware(role string) gin.HandlerFunc {
	return func (c *gin.Context) {
		roleval, exist := c.Get("role")

		if !exist {
			c.JSON(403, gin.H{"error": "role not found"})
			c.Abort()
			return 
		}

		if roleval != role {
			c.JSON(403, gin.H{"error": "insufficient permissions", })
			c.Abort()
			return 
		}

		c.Next()
	}
}