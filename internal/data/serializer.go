package data

import (
	"strings"
)

// GetUsers gets jwt users.
func GetUsers(input string) []string {
	parts := strings.Split(input, ",")

	var trimmedParts []string
	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		if len(trimmedPart) > 0 {
			trimmedParts = append(trimmedParts, trimmedPart)
		}
	}

	return trimmedParts
}
