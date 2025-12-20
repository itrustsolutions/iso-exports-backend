package businesserrors

import "fmt"

type BusinessError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
	Err     error                  `json:"-"` // Underlying error, not exposed in JSON response
}

func (e *BusinessError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (underlying: %v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *BusinessError) Unwrap() error {
	return e.Err
}

func NewBusinessError(code string, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
		Details: make(map[string]interface{}),
	}
}

func NewValidationError(errors map[string][]string) *BusinessError {
	// Convert map[string][]string to map[string]interface{}
	details := make(map[string]interface{})
	for key, value := range errors {
		details[key] = value
	}

	return NewBusinessError(ErrCodeInvalidJSONInput, "Invalid input").
		WithDetails(details)
}

func (e *BusinessError) WithDetails(details map[string]interface{}) *BusinessError {
	e.Details = details
	return e
}

func (e *BusinessError) WithError(err error) *BusinessError {
	e.Err = err
	return e
}
