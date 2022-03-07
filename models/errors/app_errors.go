package app_errors

import "errors"

var (
	ALREADY_IN_COUPLE = errors.New("user is already in a couple")
)

type AppError struct {
	Error string `json:"error"`
}
