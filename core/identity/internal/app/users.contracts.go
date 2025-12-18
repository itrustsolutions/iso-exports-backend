package app

import "context"

type UsersAppContract interface {
	CreateUser(ctx context.Context, input *CreateUserInput) (*CreateUserResult, error)
}
