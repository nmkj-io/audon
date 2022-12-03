package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

var (
	ErrInvalidRequestFormat = echo.NewHTTPError(http.StatusBadRequest, "invalid_request_format")
	ErrInvalidCookie        = echo.NewHTTPError(http.StatusBadRequest, err_invalid_cookie.Error())
	ErrSessionNotAvailable  = echo.NewHTTPError(http.StatusInternalServerError, "session_not_available")
)

func wrapValidationError(err error) error {
	wrapped := errors.Wrap(err, "validation_failed")
	return echo.NewHTTPError(http.StatusBadRequest, wrapped.Error())
}
