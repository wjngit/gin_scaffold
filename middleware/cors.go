package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	//"net/http"
)

// 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			//c.AbortWithStatus(http.StatusNoContent)
			ResponseError(c, InternalErrorCode, errors.New("invalid option"))
			c.Abort()
			return
		}
		c.Next()
	}
}
