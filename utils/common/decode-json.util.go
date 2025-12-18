package common

import (
	"encoding/json"
	"io"

	businesserrors "github.com/itrustsolutions/iso-exports-backend/utils/errors/business"
)

func DecodeJSON(body io.ReadCloser, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return businesserrors.NewBusinessError(businesserrors.ErrCodeInvalidJSONInput, "Ops! Invalid input provided.").WithError(err)
	}

	return nil
}
