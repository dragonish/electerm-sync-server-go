package logger

import (
	"electerm/internal/model"
	"io"
	"log/slog"
	"os"

	goLogger "github.com/donnie4w/go-logger/logger"
)

// instance is default logger instance.
var instance *slog.Logger

func init() {
	//? Do not initialize logs during testing.
	if os.Getenv("LOGGING_TEST") == "true" {
		testH := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
		instance = slog.New(testH)
		slog.SetDefault(instance)
		return
	}

	logFile, err := goLogger.NewLogger().SetRollingDaily(model.LOGS_DIR, model.LOG_FILENAME)
	if err != nil {
		slog.New(slog.NewTextHandler(os.Stdout, nil)).Error("unable to create log file", "error", err)
		os.Exit(1)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	h := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelInfo})
	instance = slog.New(h)

	// Set as default logger.
	slog.SetDefault(instance)
}
