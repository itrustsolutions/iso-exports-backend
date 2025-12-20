package common

import (
	"encoding/json"
	"io"
)

func DecodeJSON(body io.ReadCloser, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}
