package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

// decodes json request body into struct
func decodeJson[T any](body io.Reader) (payload *T, err error) {
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&payload)

	if err != nil {
		return nil, err
	}

	return payload, nil
}

// validates json request body struct
// returns an error message if not ok
func validateJson[T any](payload *T) (ok bool, msg []byte, err error) {
	// validate
	err = validate.Struct(payload)
	if err != nil {
		// generate error message
		m := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			m[err.Field()] = fmt.Sprintf("%s,%s", err.Tag(), err.Type())
		}
		j, err := json.Marshal(m)
		if err != nil {
			return false, make([]byte, 0), err
		}
		return false, j, err
	}
	return true, make([]byte, 0), nil
}
