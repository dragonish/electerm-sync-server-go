package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"electerm/internal/logger"
)

// ok makes a response to successful data.
func ok(ctx *gin.Context, resData interface{}) {
	ctx.JSON(http.StatusOK, resData)
}

// okWithFile makes a file response to successful data.
func okWithFile(ctx *gin.Context, filepath string) {
	ctx.Status(http.StatusOK)
	ctx.File(filepath)
}

// badRequest makes specific request error response.
func badRequest(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, errorResponse(msg))
}

// badRequestWithParse makes a response that cannot parse the request.
func badRequestWithParse(ctx *gin.Context, err error) {
	logger.WarnWithErr("parse request error", err, "path", parseRequestPath(ctx.Request))
	badRequest(ctx, "Unable to parse request")
}

// notFound makes a response with record not found.
func notFound(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusNotFound, errorResponse(msg))
}

// internalServerError makes an internal error response.
func internalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, errorResponse(logger.ErrMsg(err)))
}

// errorResponse generates body for error response.
func errorResponse(msg string) gin.H {
	return gin.H{
		"message": msg,
	}
}

// parseRequestPath returns the path of the request.
func parseRequestPath(r *http.Request) string {
	return r.URL.Path
}
