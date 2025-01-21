package apis

import (
	"electerm/internal/global"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		claims := jwt.ExtractClaims(c)
		id := claims["id"].(string)
		return global.Flags.Includes(id)
	}
}
