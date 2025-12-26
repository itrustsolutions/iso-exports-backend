package db

import (
	"context"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	technicalerrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/technical"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Executes a unit of work within a managed PostgreSQL transaction.
//
// It automates the transaction lifecycle:
//  1. Begins a transaction from the provided pool.
//  2. Injects the transaction into the context using customcontext.WithTx.
//  3. Executes the provided 'fn' closure.
//  4. Commits on success or rolls back on error/panic.
//
// If 'fn' returns an error, that specific error is returned to the caller
// after rollback. If the Begin or Commit stages fail, a TechnicalError
// is returned with the appropriate error code.
//
// This function is panic-safe; it will trigger a rollback before re-panicking
// to ensure no database locks are held open.
func ExecWithinTx[T any](ctx context.Context, pool *pgxpool.Pool, fn func(ctx context.Context) (T, error)) (T, error) {
	var result T

	// 1. Start the transaction
	tx, err := pool.Begin(ctx)
	if err != nil {
		return result, technicalerrors.NewTechnicalError(
			technicalerrors.ErrCodeTransactionFailedToBegin,
			"Failed to begin transaction",
		).WithError(err)
	}

	// 2. Ensure rollback on panic
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	// 3. Inject tx into context
	txCtx := customcontext.WithTx(ctx, tx)

	// 4. Run the logic
	result, err = fn(txCtx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return result, err
	}

	// 5. Commit
	if err := tx.Commit(ctx); err != nil {
		return result, technicalerrors.NewTechnicalError(
			technicalerrors.ErrCodeTransactionFailedToCommit,
			"Failed to commit transaction",
		).WithError(err)
	}

	return result, nil
}
