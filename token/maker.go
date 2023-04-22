package token

import (
	"time"
)

// Maker is an interface for managing tokens
type Maker interface {
	CreateToken(id int64, duration time.Duration) (string, error)

	// VerifyToken checks if the tooken is valid or not
	VerifyToken(token string) (*Payload, error)
}
