package security

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// NewID generates a new ULID and returns it as a 26-character string.
func NewID() string {
	// Current timestamp in UTC
	timestamp := time.Now().UTC()

	// Entropy source for randomness, monotonic ensures ordering if called quickly
	entropy := ulid.Monotonic(rand.Reader, 0)

	// Generate ULID
	id := ulid.MustNew(ulid.Timestamp(timestamp), entropy)

	return id.String()
}
