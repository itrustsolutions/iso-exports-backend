package technicalerrors

import (
	"fmt"
	"net/http"
)

type TechnicalError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Details    map[string]interface{} `json:"details,omitempty"`
	Err        error                  `json:"-"` // Underlying error, not exposed in JSON response
	StatusCode int                    `json:"-"` // HTTP status code, not exposed in JSON response. Default to http.StatusInternalServerError
}

func NewTechnicalError(code string, message string) *TechnicalError {
	return &TechnicalError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func (e *TechnicalError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (underlying: %v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *TechnicalError) Unwrap() error {
	return e.Err
}

func (e *TechnicalError) WithDetails(details map[string]interface{}) *TechnicalError {
	e.Details = details
	return e
}

func (e *TechnicalError) WithError(err error) *TechnicalError {
	e.Err = err
	return e
}

func (e *TechnicalError) WithHTTPStatus(status int) *TechnicalError {
	e.StatusCode = status
	return e
}

func AsTechnicalError(err error, target **TechnicalError) bool {
	if te, ok := err.(*TechnicalError); ok {
		*target = te
		return true
	}
	return false
}
