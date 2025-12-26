package customcontext

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TxKey struct{}

// Extracts the transaction from the context if exists or return nil
func ExtractTx(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}

func WithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey{}, tx)
}
