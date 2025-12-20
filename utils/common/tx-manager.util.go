package common

import (
	"context"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TXManager struct {
	pool *pgxpool.Pool
}

func NewTXManager(pool *pgxpool.Pool) *TXManager {
	return &TXManager{pool: pool}
}

// Begin returns a new context with a transaction attached, and the tx itself for manual control
func (p *TXManager) Begin(ctx context.Context) (context.Context, pgx.Tx, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Inject the transaction into the context
	newCtx := context.WithValue(ctx, customcontext.TxKey{}, tx)

	return newCtx, tx, nil
}
