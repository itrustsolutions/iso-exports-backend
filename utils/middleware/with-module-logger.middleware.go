package middleware

import (
	"net/http"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
)

// A middleware that adds a module-specific logger to the request context
func WithModuleLogger(moduleName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract existing logger from context
			logger := customcontext.ExtractLogger(r.Context())

			// Create a new logger with module field
			newLogger := logger.With().Str("module", moduleName).Logger()
			logger = &newLogger

			// Store updated logger back in context
			ctx := customcontext.WithLogger(r.Context(), logger)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
