package httppresentation

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/app"
	identitydtos "github.com/itrustsolutions/iso-exports-backend/core/identity/pkg/dtos"
	"github.com/itrustsolutions/iso-exports-backend/utils/db"
	httputils "github.com/itrustsolutions/iso-exports-backend/utils/http"
	"github.com/itrustsolutions/iso-exports-backend/utils/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRouter struct {
	usersApplication *app.UsersApp
	pool             *pgxpool.Pool
	r                *chi.Mux
}

func NewUsersRouter(usersApp *app.UsersApp, pool *pgxpool.Pool, r *chi.Mux) *UsersRouter {
	return &UsersRouter{
		usersApplication: usersApp,
		pool:             pool,
		r:                r,
	}
}

func (ur *UsersRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", middleware.MakeHandler(ur.postCreateUserHandler))

	return r
}

func (ur *UsersRouter) postCreateUserHandler(
	w http.ResponseWriter, r *http.Request,
) (*middleware.SuccessResult, error) {
	input := &identitydtos.CreateUserInput{}

	err := httputils.DecodeJSON(r.Body, input)
	if err != nil {
		return nil, err
	}

	result, err := db.ExecWithinTx(r.Context(), ur.pool, func(txCtx context.Context) (*identitydtos.CreateUserResult, error) {
		return ur.usersApplication.CreateUser(txCtx, input)
	})

	if err != nil {
		return nil, err
	}

	return middleware.NewSuccessResult(result, http.StatusCreated), nil
}
