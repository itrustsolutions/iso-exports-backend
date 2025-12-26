package identity

import (
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/app"
	db "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/db/gen"
	"github.com/itrustsolutions/iso-exports-backend/core/identity/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	Users app.UsersAppContract
}

type Config struct {
	DB *pgxpool.Pool
}

func NewModule(cfg *Config) *Module {
	queries := db.New(cfg.DB)

	usersService := domain.NewUsersService(queries)
	usersApp := app.NewUsersApp(usersService)

	return &Module{
		Users: usersApp,
	}
}
