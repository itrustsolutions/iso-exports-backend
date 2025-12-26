package httputils

import (
	"encoding/json"
	"io"

	businesserrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/business"
)

// Decodes a JSON payload from an io.ReadCloser into the provided interface v.
//
// It disallows unknown fields and returns a business error if decoding fails.
func DecodeJSON(body io.ReadCloser, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return businesserrors.NewBusinessError(
			businesserrors.ErrCodeInvalidJSONInput,
			"Invalid JSON input",
		)
	}

	return nil
}
