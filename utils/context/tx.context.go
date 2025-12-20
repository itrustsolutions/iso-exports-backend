package customcontext

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TxKey struct{}

// Extract retrieves the transaction or nil if none exists
func ExtractTx(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}
