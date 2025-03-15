package parse

import (
	"electerm/internal/logger"
	"electerm/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

// ParseJWT parses JWT string.
func ParseJWT(secret string, tokenString string) (model.TokenClaims, error) {
	var resClaims model.TokenClaims
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		logger.WarnWithErr("the jwt invalid", err)
		return resClaims, err
	}

	if claims, ok := token.Claims.(*model.JWTClaims); ok {
		resClaims = claims.TokenClaims
	}

	return resClaims, nil
}
