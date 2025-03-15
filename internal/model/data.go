package model

import "github.com/golang-jwt/jwt/v5"

// DATA_DIR defines the data file directory.
const DATA_DIR = "data"

// TokenClaims defines token claims.
type TokenClaims struct {
	ID string `json:"id"`
}

// JWTClaims defines JWT claims.
type JWTClaims struct {
	TokenClaims
	jwt.RegisteredClaims
}
