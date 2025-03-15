package apis

import (
	"electerm/internal/store"

	"github.com/gin-gonic/gin"
)

func Apis(router *gin.Engine) {
	authGroup := router.Group("/api", JWTMiddleware())

	authGroup.PUT("/sync", func(ctx *gin.Context) {
		id := ctx.GetString("id")
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
		id := ctx.GetString("id")
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
