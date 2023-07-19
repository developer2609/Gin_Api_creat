package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Token not Present",
		})
		return

	}
	c.Next()

}
