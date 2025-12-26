package middleware

import (
	"encoding/json"
	"net/http"

	httputils "github.com/itrustsolutions/iso-exports-backend/utils/http"
)

type SuccessResult struct {
	data   interface{}
	status int
}

func NewSuccessResult(data interface{}, status int) *SuccessResult {
	return &SuccessResult{
		data:   data,
		status: status,
	}
}

type AppHTTPHandler func(w http.ResponseWriter, r *http.Request) (*SuccessResult, error)

func MakeHandler(handler AppHTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := handler(w, r)

		if err != nil {
			errResponse := httputils.NewErrorResponse(r.Context(), err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(errResponse.Status)
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		successResponse := httputils.NewSuccessResponse(r.Context(), result.data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(result.status)
		json.NewEncoder(w).Encode(successResponse)
	}
}
