package logger

import "github.com/gin-gonic/gin"

// LoggerMiddleware returns a gin.HandlerFunc for logging request logs
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		Info("request", "method", ctx.Request.Method, "path", ctx.Request.URL.Path, "status", ctx.Writer.Status())
	}
}
