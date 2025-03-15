package model

import "strconv"

// Flags serves global state to the program.
type Flags struct {
	Port      uint16   // web service running port.
	JWTSecret string   // JWT secret
	JWTUsers  []string // JWT users
}

// GetPortStr returns port string.
func (f *Flags) GetPortStr() string {
	return strconv.FormatUint(uint64(f.Port), 10)
}

// Includes returns user's existence status.
func (f *Flags) Includes(user string) bool {
	for _, item := range f.JWTUsers {
		if item == user {
			return true
		}
	}
	return false
}
