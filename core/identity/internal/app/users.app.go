package app

import (
	"context"

	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/domain"
	identitydtos "github.com/itrustsolutions/iso-exports-backend/core/identity/pkg/dtos"
	"github.com/itrustsolutions/iso-exports-backend/utils/common"
)

type UsersApp struct {
	tXManager    *common.TXManager
	usersService *domain.UsersService
}

func NewUsersApp(usersService *domain.UsersService) *UsersApp {
	return &UsersApp{
		usersService: usersService,
	}
}

func (a *UsersApp) CreateUser(ctx context.Context, input *identitydtos.CreateUserInput) (*identitydtos.CreateUserResult, error) {
	err := input.Validate()

	if err != nil {
		return nil, err
	}

	result, err := a.usersService.CreateUser(ctx, &domain.CreateUserInput{
		Username:        input.Username,
		Email:           input.Email,
		PlainPassword:   input.PlainPassword,
		HasSystemAccess: input.HasSystemAccess,
		IsActive:        input.IsActive,
	})

	if err != nil {
		return nil, err
	}

	return &identitydtos.CreateUserResult{
		ID:                     result.ID,
		Username:               result.Username,
		Email:                  result.Email,
		HasSystemAccess:        result.HasSystemAccess,
		HasAllNamespacesAccess: result.HasAllNamespacesAccess,
		IsActive:               result.IsActive,
		CreatedAt:              result.CreatedAt,
		UpdatedAt:              result.UpdatedAt,
	}, nil
}
