package domain

import "context"

type UsersServiceContract interface {
	CreateUser(ctx context.Context, input *CreateUserInput) (*CreateUserResult, error)
}
