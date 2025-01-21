package parse

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnvVars(t *testing.T) {
	envs := parseEnvVars()
	assert.Equal(t, uint16(7837), envs.Port)
	assert.Empty(t, envs.JWTSecret)
	assert.Empty(t, envs.JWTUsers)

	os.Setenv("ELECTERM_PORT", "8080")
	os.Setenv("ELECTERM_JWT_SECRET", "abcdefg")
	os.Setenv("ELECTERM_JWT_USERS", "user1,user2,user3")

	envs = parseEnvVars()
	assert.Equal(t, uint16(8080), envs.Port)
	assert.Equal(t, "abcdefg", envs.JWTSecret)
	assert.Equal(t, []string{"user1", "user2", "user3"}, envs.JWTUsers)
}
