package parse

import (
	"electerm/internal/data"
	"electerm/internal/global"
	"electerm/internal/logger"
)

// Parse handles environment variables.
//
// Assign values to the global flag state.
func Parse() {
	flags := parseEnvVars()

	if len(flags.JWTSecret) == 0 {
		logger.Fatal("incorrect environment variable", logger.NewErr("jwt secret is empty"))
	}

	if len(flags.JWTUsers) == 0 {
		logger.Fatal("incorrect environment variable", logger.NewErr("jwt users is empty"))
	}

	logger.Info("web service running port", "port", flags.Port)
	logger.Info("jwt secret", "secret", data.MaskWithStars(flags.JWTSecret))
	logger.Info("jwt users", "users", flags.JWTUsers)

	global.Flags = flags
}
