package domain

import (
	"time"
)

type CreateUserInput struct {
	Username        string
	Email           string
	PlainPassword   string
	HasSystemAccess bool
	IsActive        bool
}

type CreateUserResult struct {
	ID                     string
	Username               string
	Email                  string
	HasSystemAccess        bool
	HasAllNamespacesAccess bool
	IsActive               bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
