package utils

import (
	"errors"
	"strings"
)

func CleanInput(input string) ([]string, error) {
	if input == "" {
		return nil, errors.New("input is empty")
	}

	lower := strings.ToLower(input)
	trimmed := strings.Fields(lower)
	return trimmed, nil
}
