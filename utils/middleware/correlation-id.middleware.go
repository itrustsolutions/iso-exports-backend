package middleware

import (
	"net/http"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	"github.com/itrustsolutions/iso-exports-backend/utils/security"
)

// A middleware that assigns a correlation id to the context.
//
// If the `X-Correlation-ID` header exists with a value it will be set in the context.
// Otherwise, a new correlation id will be generated and set in the context.
func CorrelationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = security.NewID()
		}

		ctx := customcontext.WithCorrelationId(r.Context(), correlationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
