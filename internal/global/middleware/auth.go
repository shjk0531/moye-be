package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: JWT 토큰 검증 등 인증 로직 구현
		c.Next()
	}
}
