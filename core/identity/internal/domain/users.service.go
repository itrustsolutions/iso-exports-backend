package domain

import (
	"context"

	db "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/db/gen"
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/db/utils"
	errorscodes "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/errors"
	businesserrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/business"
	customdberrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/db"
	"github.com/itrustsolutions/iso-exports-backend/utils/errors/technical"
	"github.com/itrustsolutions/iso-exports-backend/utils/security"
)

type UsersService struct {
	queries *db.Queries
}

func NewUsersService(queries *db.Queries) *UsersService {
	return &UsersService{
		queries: queries,
	}
}

func (s *UsersService) CreateUser(ctx context.Context, input *CreateUserInput) (*CreateUserResult, error) {
	queries := utils.GetQueriesWithPossibleTx(ctx, s.queries)

	hashedPassword, err := security.HashString(input.PlainPassword)
	if err != nil {
		return nil, technical.NewTechnicalError(
			technical.ErrCodeHashingPasswordFailed,
			"Failed to hash user password",
		).WithError(err)
	}

	newID := security.NewID()

	result, err := queries.CreateUser(ctx, db.CreateUserParams{
		ID:              newID,
		Username:        input.Username,
		Email:           input.Email,
		PasswordHash:    hashedPassword,
		HasSystemAccess: input.HasSystemAccess,
		IsActive:        input.IsActive,
	})

	if err != nil {
		pgErr := customdberrors.GetPgError(err)

		if pgErr != nil {
			switch customdberrors.MapPostgresError(err) {
			case customdberrors.ErrDBUniqueViolation:
				switch pgErr.ConstraintName {
				case "uq_users_username":
					return nil, businesserrors.NewBusinessError(
						errorscodes.ErrCodeUserWithUsernameExists,
						"Username already exists",
					).WithError(err).WithDetails(map[string]interface{}{
						"username": input.Username,
					})
				case "uq_users_email":
					return nil, businesserrors.NewBusinessError(
						errorscodes.ErrCodeUserWithEmailExists,
						"Email already exists",
					).WithError(err).WithDetails(map[string]interface{}{
						"email": input.Email,
					})
				}
			}
		}

		// Non-DB error return a generic internal error
		return nil, technical.NewTechnicalError(
			technical.ErrCodeDatabaseOperationFailed,
			"Failed to create user in database",
		).WithError(err)
	}

	return &CreateUserResult{
		ID:                     result.ID,
		Username:               result.Username,
		Email:                  result.Email,
		HasSystemAccess:        result.HasSystemAccess,
		HasAllNamespacesAccess: result.HasAllNamespacesAccess,
		IsActive:               result.IsActive,
		CreatedAt:              result.CreatedAt.Time,
		UpdatedAt:              result.UpdatedAt.Time,
	}, nil
}

var _ UsersServiceContract = (*UsersService)(nil)
