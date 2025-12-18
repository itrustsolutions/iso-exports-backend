package app

import (
	"time"

	"github.com/asaskevich/govalidator"
	businesserrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/business"
)

type CreateUserInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	PlainPassword   string `json:"plainPassword"`
	HasSystemAccess bool   `json:"hasSystemAccess"`
	IsActive        bool   `json:"isActive"`
}

func (input *CreateUserInput) Validate() error {
	errors := make(map[string][]string)

	// TODO: Fix the validation rules those just samples

	// Username validation
	if govalidator.IsNull(input.Username) {
		errors["username"] = append(errors["username"], "Username is required")
	} else if len(input.Username) < 3 || len(input.Username) > 50 {
		errors["username"] = append(errors["username"], "Username must be between 3 and 50 characters")
	}

	// Email validation
	if govalidator.IsNull(input.Email) {
		errors["email"] = append(errors["email"], "Email is required")
	} else if !govalidator.IsEmail(input.Email) {
		errors["email"] = append(errors["email"], "Email format is invalid")
	}

	// Password validation
	if govalidator.IsNull(input.PlainPassword) {
		errors["plainPassword"] = append(errors["plainPassword"], "Password is required")
	} else if len(input.PlainPassword) < 6 {
		errors["plainPassword"] = append(errors["plainPassword"], "Password must be at least 6 characters long")
	}

	if len(errors) > 0 {
		return businesserrors.NewValidationError(errors)
	}

	return nil
}

type CreateUserResult struct {
	ID                     string    `json:"id"`
	Username               string    `json:"username"`
	Email                  string    `json:"email"`
	HasSystemAccess        bool      `json:"hasSystemAccess"`
	HasAllNamespacesAccess bool      `json:"hasAllNamespacesAccess"`
	IsActive               bool      `json:"isActive"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
