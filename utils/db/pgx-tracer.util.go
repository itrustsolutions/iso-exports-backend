package db

import (
	"context"
	"strings"
	"time"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	"github.com/jackc/pgx/v5"
)

// PgxQueryTracer logs pgx query events using the request-scoped zerolog from context.
// It captures SQL, args, duration, and success/failure, tied to the correlation ID.
type PgxQueryTracer struct{}

func NewPgxQueryTracer() *PgxQueryTracer {
	return &PgxQueryTracer{}
}

// TraceQueryStart records the query start time and returns an augmented context.
func (t *PgxQueryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	logger := customcontext.ExtractLogger(ctx)
	if logger != nil {
		evt := logger.Debug().
			Str("sql", cleanSQL(data.SQL))
		// Only log args if not nil (avoid args=null)
		if data.Args != nil {
			evt = evt.Interface("args", data.Args)
		}
		evt.Msg("db.query.start")
	}

	return customcontext.WithPgxTracerCtxData(ctx, customcontext.PgxTracerCtxData{
		Sql: data.SQL, Args: data.Args, Started: time.Now(),
	})
}

// TraceQueryEnd logs the query execution with duration and outcome.
func (t *PgxQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	logger := customcontext.ExtractLogger(ctx)
	if logger == nil {
		return
	}

	d := customcontext.ExtractPgxTracerCtxData(ctx)
	if d == nil {
		d = &customcontext.PgxTracerCtxData{}
	}

	duration := time.Since(d.Started)

	if data.Err != nil {
		evt := logger.Error().
			Str("sql", cleanSQL(d.Sql)).
			Int64("duration_ms", duration.Milliseconds()).
			Err(data.Err)
		// Only log args if not nil
		if d.Args != nil {
			evt = evt.Interface("args", d.Args)
		}
		evt.Msg("db.query.failed")
		return
	}
	evt := logger.Info().
		Str("sql", cleanSQL(d.Sql)).
		Int64("duration_ms", duration.Milliseconds())
	// Only log args if not nil
	if d.Args != nil {
		evt = evt.Interface("args", d.Args)
	}
	evt.Msg("db.query.succeeded")
}

// cleanSQL removes newlines and excessive whitespace from SQL statements for cleaner logging.
func cleanSQL(sql string) string {
	// Replace newlines and tabs with spaces
	sql = strings.ReplaceAll(sql, "\n", " ")
	sql = strings.ReplaceAll(sql, "\t", " ")

	// Replace multiple spaces with single space
	sql = strings.Join(strings.Fields(sql), " ")

	return strings.TrimSpace(sql)
}
