package parse

import (
	"encoding/json"
	"errors"

	"../structs"
	"github.com/google/uuid"
)

// Validate validates an incoming JSON request body. It returns a CompileRequest on success
// and an Error on validation or JSON parsing error.
func Validate(raw []byte) (structs.CompileRequest, error) {
	var cr structs.CompileRequest
	err := json.Unmarshal(raw, &cr)
	if err != nil {
		return cr, err
	}

	if cr.ConfigHeader == "" {
		return cr, errors.New("config_header must not be empty")
	}
	if cr.PioConfig == "" {
		return cr, errors.New("pio_config must not be empty")
	}

	cr.ID = uuid.New().String()
	return cr, nil
}
