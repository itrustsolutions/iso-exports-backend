package app

import (
	"context"

	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/domain"
	"github.com/itrustsolutions/iso-exports-backend/utils/common"
)

type UsersApp struct {
	tXManager    *common.TXManager
	usersService *domain.UsersService
}

func NewUsersApp(usersService *domain.UsersService, tXManager *common.TXManager) *UsersApp {
	return &UsersApp{
		tXManager:    tXManager,
		usersService: usersService,
	}
}

func (a *UsersApp) CreateUser(ctx context.Context, input *CreateUserInput) (*CreateUserResult, error) {
	err := input.Validate()

	if err != nil {
		return nil, err
	}

	txCtx, tx, err := a.tXManager.Begin(ctx)
	defer tx.Rollback(txCtx)

	result, err := a.usersService.CreateUser(txCtx, &domain.CreateUserInput{
		Username:        input.Username,
		Email:           input.Email,
		PlainPassword:   input.PlainPassword,
		HasSystemAccess: input.HasSystemAccess,
		IsActive:        input.IsActive,
	})

	if err != nil {
		return nil, err
	}

	return &CreateUserResult{
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
