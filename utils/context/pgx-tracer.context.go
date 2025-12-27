package customcontext

import (
	"context"
	"time"
)

type PgxTracerCtxKey struct{}

type PgxTracerCtxData struct {
	Sql     string
	Args    []any
	Started time.Time
}

func ExtractPgxTracerCtxData(ctx context.Context) *PgxTracerCtxData {
	if v := ctx.Value(PgxTracerCtxKey{}); v != nil {
		if td, ok := v.(PgxTracerCtxData); ok {
			return &td
		}
	}
	return nil
}

func WithPgxTracerCtxData(ctx context.Context, data PgxTracerCtxData) context.Context {
	return context.WithValue(ctx, PgxTracerCtxKey{}, data)
}
