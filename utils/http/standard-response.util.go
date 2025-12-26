package httputils

import (
	"context"
	"net/http"

	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	businesserrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/business"
	technicalerrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/technical"
)

type StandardResponse struct {
	Ok            bool           `json:"ok"`
	Status        int            `json:"status"`
	CorrelationID string         `json:"correlationId,omitempty"`
	Data          interface{}    `json:"data,omitempty"`
	Error         *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func NewErrorResponse(ctx context.Context, err error) *StandardResponse {
	correlationId := customcontext.ExtractCorrelationId(ctx)

	var code, message string
	var details map[string]interface{}
	var status int

	// Business error
	var be *businesserrors.BusinessError
	if ok := businesserrors.AsBusinessError(err, &be); ok {
		code = be.Code
		message = be.Message
		details = be.Details
		status = be.StatusCode
		if status == 0 {
			status = http.StatusBadRequest
		}
	} else {
		// Technical error
		var te *technicalerrors.TechnicalError
		if ok := technicalerrors.AsTechnicalError(err, &te); ok {
			code = te.Code
			message = te.Message
			details = te.Details
			status = te.StatusCode
			if status == 0 {
				status = http.StatusInternalServerError
			}
		} else {
			// Fallback for unknown errors
			code = "INTERNAL_SERVER_ERROR"
			message = "An unexpected error occurred"
			details = map[string]interface{}{"error": err.Error()}
			status = http.StatusInternalServerError
		}
	}

	return &StandardResponse{
		Ok:            false,
		Status:        status,
		CorrelationID: correlationId,
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
}

func NewSuccessResponse(ctx context.Context, data interface{}) *StandardResponse {
	correlationId := customcontext.ExtractCorrelationId(ctx)

	return &StandardResponse{
		Ok:            true,
		Status:        http.StatusOK,
		CorrelationID: correlationId,
		Data:          data,
	}
}
