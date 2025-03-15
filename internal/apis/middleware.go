package apis

import (
	"electerm/internal/global"
	"electerm/internal/logger"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
)

func InitParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Key:         []byte(global.Flags.JWTSecret),
		Timeout:     time.Minute * 15,
		IdentityKey: "id",

		Authorizator:  authorizator(),
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc: func() time.Time {
			return time.Now().Add(time.Second * time.Duration(global.Flags.TimeError))
		},
	}
}

// JWTMiddleware returns jwt middleware.
func JWTMiddleware(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		if errInit := authMiddleware.MiddlewareInit(); errInit != nil {
			logger.Err("middleware init error", errInit)
		}
	}
}
