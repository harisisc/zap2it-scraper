package zap2it

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)
