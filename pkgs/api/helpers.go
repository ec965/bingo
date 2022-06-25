package api

import (
	"encoding/json"
	"errors"
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
func validatePayload[T any](payload *T) error {
	// validate
	err := validate.Struct(payload)
	if err != nil {
		// generate error message
		msg := ""
		for _, err := range err.(validator.ValidationErrors) {
			msg += fmt.Sprintf("%s:%s,%s;", err.Field(), err.Tag(), err.Type())
		}
		return errors.New(msg)
	}
	return nil
}
