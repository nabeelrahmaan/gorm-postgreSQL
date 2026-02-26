package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		store := cookie.NewStore([]byte("secret-key"))
		sessions.Sessions("mysession", store)
	})
}