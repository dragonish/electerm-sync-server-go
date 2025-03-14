package apis

import (
	"electerm/internal/store"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Apis(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	authGroup := router.Group("/api", authMiddleware.MiddlewareFunc())

	authGroup.PUT("/sync", func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		id := claims["id"].(string)
		var body interface{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			badRequestWithParse(ctx, err)
			return
		}

		if err := store.Writer(id, body); err != nil {
			internalServerError(ctx, err)
		} else {
			ok(ctx, "ok")
		}
	})

	authGroup.POST("/sync", func(ctx *gin.Context) {
		ok(ctx, "ok")
	})

	authGroup.GET("/sync", func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		id := claims["id"].(string)

		filepath := store.Reader(id)
		if len(filepath) > 0 {
			okWithJSONFile(ctx, filepath)
		} else {
			notFound(ctx, "not found")
		}
	})

	router.GET("/test", func(ctx *gin.Context) {
		ok(ctx, "ok")
	})

	router.NoRoute(func(ctx *gin.Context) {
		notFound(ctx, "not found")
	})
}
