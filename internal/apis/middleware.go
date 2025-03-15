package apis

import (
	"electerm/internal/global"
	"electerm/internal/parse"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware returns jwt middleware.
func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("Authorization")
		if header == "" {
			context.JSON(401, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}
		tokenString := strings.Split(header, " ")[1]
		token, err := parse.ParseJWT(global.Flags.JWTSecret, tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}
		if global.Flags.Includes(token.ID) {
			context.Set("id", token.ID)
			context.Next()
		} else {
			context.JSON(401, gin.H{"error": "Unauthorized"})
			context.Abort()
		}
	}
}
