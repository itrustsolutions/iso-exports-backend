package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	httputils "github.com/itrustsolutions/iso-exports-backend/utils/http"
	"github.com/rs/zerolog"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// **Note**: This middleware assumes that CorrelationID middleware has already been applied.
func RequestLoggingMiddleware(baseLogger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create request-scoped logger with correlation ID
			logger := baseLogger.With().
				Str("correlation_id", customcontext.ExtractCorrelationId(r.Context())).
				Str("replica_id", fmt.Sprintf("%d", os.Getpid())).
				Logger()

			// Log request start with all HTTP context
			logger.Info().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Str("ip", httputils.GetClientIP(r)).
				Str("user_agent", r.UserAgent()).
				// These would come from auth middleware
				// TODO: integrate with auth to get real session/user IDs
				Str("session_id", ""). // placeholder
				Str("user_id", "").    // placeholder
				Msg("request started")

			ctx := customcontext.WithLogger(r.Context(), &logger)

			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(ww, r.WithContext(ctx))

			duration := time.Since(start).Milliseconds()
			logger.Info().
				Int("status", ww.statusCode).
				Int64("duration_ms", duration).
				Msg("request completed")
		})
	}
}
