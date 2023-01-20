package Validator

import (
	"errors"
)

func Not_empty_string(input string) error {
	if input == "" {
		return errors.New("Invalid empty string")
	}
	return nil
}
