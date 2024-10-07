package middleware

import (
	"src/index/src/utils"

	"github.com/gin-gonic/gin"
)

func NetworkLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Logger.Infof("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
