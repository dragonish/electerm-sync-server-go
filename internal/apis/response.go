package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"electerm/internal/logger"
)

// ok makes a response to successful data.
func ok(ctx *gin.Context, resData string) {
	ctx.String(http.StatusOK, resData)
}

// okWithJSONFile makes a file response to successful data.
func okWithJSONFile(ctx *gin.Context, filepath string) {
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", "application/json")
	ctx.File(filepath)
}

// badRequest makes specific request error response.
func badRequest(ctx *gin.Context, msg string) {
	ctx.String(http.StatusBadRequest, msg)
}

// badRequestWithParse makes a response that cannot parse the request.
func badRequestWithParse(ctx *gin.Context, err error) {
	logger.WarnWithErr("parse request error", err, "path", parseRequestPath(ctx.Request))
	badRequest(ctx, "Unable to parse request")
}

// notFound makes a response with record not found.
func notFound(ctx *gin.Context, msg string) {
	ctx.String(http.StatusNotFound, msg)
}

// internalServerError makes an internal error response.
func internalServerError(ctx *gin.Context, err error) {
	ctx.String(http.StatusInternalServerError, logger.ErrMsg(err))
}

// parseRequestPath returns the path of the request.
func parseRequestPath(r *http.Request) string {
	return r.URL.Path
}
