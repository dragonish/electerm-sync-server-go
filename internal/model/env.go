package model

// ENV_VAR_PREFIX is environment variable name prefix.
const ENV_VAR_PREFIX = "ELECTERM_"

// Env storages environment variable.
type Env struct {
	Port      uint16 `env:"PORT" envDefault:"7837"`    // web service running port. default: 7837
	JWTSecret string `env:"JWT_SECRET" envDefault:""`  // JWT secret
	JWTUsers  string `env:"JWT_USERS" envDefault:""`   // JWT users
	TimeError int8   `env:"TIME_ERROR" envDefault:"0"` // time error tolerance, unit: second. default: 0
}
