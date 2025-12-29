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
	logger := customcontext.ExtractLogger(ctx)

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

		// Log business error as warning
		if logger != nil {
			logger.Warn().
				Str("error_type", "business_error").
				Str("error_code", code).
				Str("error_message", message).
				Int("status", status).
				Interface("details", details).
				Msg("business error occurred")
		}
	} else {
		// Technical error
		var te *technicalerrors.TechnicalError
		if technicalerrors.AsTechnicalError(err, &te) {
			// Always show generic message & don't expose details, but respect status code
			code = technicalerrors.ErrCodeInternalServerError
			message = "Ops! Something went wrong."
			details = nil
			status = te.StatusCode
			if status == 0 {
				status = http.StatusInternalServerError
			}

			// Log technical error
			if logger != nil {
				logger.Error().
					Str("error_type", "technical_error").
					Str("error_code", te.Code). // Log actual error code for technical errors
					Str("error_message", te.Message).
					Int("status", te.StatusCode).
					Interface("details", te.Details).
					Msg("technical error occurred")
			}
		} else {
			// Unknown error: log warning, generic message
			code = technicalerrors.ErrCodeInternalServerError
			message = "Ops! Something went wrong."
			details = nil
			status = http.StatusInternalServerError

			if logger != nil {
				logger.Warn().
					Str("error_type", "unknown_error").
					Str("error_message", err.Error()).
					Int("status", status).
					Msg("unknown error occurred")
			}
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
