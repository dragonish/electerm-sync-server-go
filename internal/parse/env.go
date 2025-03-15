package parse

import (
	"electerm/internal/data"
	"electerm/internal/logger"
	"electerm/internal/model"

	"github.com/caarlos0/env/v10"
)

func parseEnvVars() (stor model.Flags) {
	cfg := model.Env{}

	opts := env.Options{
		Prefix: model.ENV_VAR_PREFIX,
	}

	if err := env.ParseWithOptions(&cfg, opts); err != err {
		logger.Err("error parsing environment variables", err)
		return
	}

	stor.Port = cfg.Port
	stor.JWTSecret = cfg.JWTSecret
	stor.JWTUsers = data.GetUsers(cfg.JWTUsers)
	stor.TimeError = cfg.TimeError

	return
}
