package identity

import (
	"github.com/go-chi/chi/v5"
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/app"
	db "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/db/gen"
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/domain"
	httppresentation "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	Users app.UsersAppContract
}

type Config struct {
	DB       *pgxpool.Pool
	Router   *chi.Mux
	HTTPPath string
}

func NewModule(cfg *Config) *Module {
	// Set up database queries (repository layer)
	queries := db.New(cfg.DB)

	// Set up domain services (resource business logic layer)
	usersService := domain.NewUsersService(queries)

	// Set up application layer (orchestration layer)
	usersApp := app.NewUsersApp(usersService)

	// Set up HTTP presentation layer (HTTP handlers)
	usersRouter := httppresentation.NewUsersRouter(usersApp, cfg.DB, cfg.Router)

	cfg.Router.Route(cfg.HTTPPath, func(r chi.Router) {
		r.Mount("/users", usersRouter.Routes())
	})

	return &Module{
		Users: usersApp,
	}
}
