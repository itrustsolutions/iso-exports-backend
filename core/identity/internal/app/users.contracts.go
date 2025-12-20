package app

import (
	"context"

	identitydtos "github.com/itrustsolutions/iso-exports-backend/core/identity/pkg/dtos"
)

type UsersAppContract interface {
	CreateUser(ctx context.Context, input *identitydtos.CreateUserInput) (*identitydtos.CreateUserResult, error)
}
