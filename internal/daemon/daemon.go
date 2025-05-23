package daemon

import (
	"context"
	"electerm/internal/apis"
	"electerm/internal/global"
	"electerm/internal/logger"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// StartDaemon starts the daemon.
func StartDaemon() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGALRM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(logger.LoggerMiddleware())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	//* Register routing
	apis.Apis(router)

	srv := &http.Server{
		Addr:              ":" + global.Flags.GetPortStr(),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Err("application startup error", err)
		}
	}()
	logger.Info("application started")

	<-ctx.Done()

	stop()
	logger.Info("application closing, press Ctrl+C to exit immediately")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Err("application is forced to close", err)
	}

	logger.Info("application has been closed")
}
