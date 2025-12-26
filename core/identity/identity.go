package identity

import (
	"net/http"

	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/app"
	db "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/db/gen"
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/domain"
	httppresentation "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	Users  app.UsersAppContract
	Routes http.Handler
}

type Config struct {
	DB *pgxpool.Pool
}

func NewModule(cfg *Config) *Module {
	queries := db.New(cfg.DB)

	usersService := domain.NewUsersService(queries)
	usersApp := app.NewUsersApp(usersService)
	usersRoutes := httppresentation.NewUsersHTTP(usersApp, cfg.DB)

	return &Module{
		Users:  usersApp,
		Routes: usersRoutes.Routes(),
	}
}
