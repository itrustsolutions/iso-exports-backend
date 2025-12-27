package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itrustsolutions/iso-exports-backend/utils/config"
	technicalerrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/technical"
	httputils "github.com/itrustsolutions/iso-exports-backend/utils/http"
)

// Recovery middleware recovers from panics and returns a JSON internal server error response.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				cfg := config.GetConfigOrExist()
				te := technicalerrors.NewTechnicalError(technicalerrors.ErrCodeInternalServerError, "Ops! Something went wrong.")

				// Only attach panic details in development or local environments
				if cfg.App.Env == "development" || cfg.App.Env == "local" {
					te = te.WithDetails(map[string]interface{}{"panic": fmt.Sprintf("%v", rec)})
				}

				errResponse := httputils.NewErrorResponse(r.Context(), te)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(errResponse.Status)
				_ = json.NewEncoder(w).Encode(errResponse)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
