package middlewares

import (
	"errors"
	"github.com/ALexfonSchneider/avito-shop/internal/application"
	"github.com/ALexfonSchneider/avito-shop/internal/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type ErrorsMiddleware struct{}

func NewErrorsMiddleware() *ErrorsMiddleware {
	return &ErrorsMiddleware{}
}

func (e *ErrorsMiddleware) httpErrorFromApplicationError(err application.Error) (resp ErrorResponse, code int) {
	code = http.StatusInternalServerError

	switch {
	case errors.Is(err, application.UserNotFound):
		code = http.StatusNotFound
	case errors.Is(err, application.NotEnoughCoins):
		code = http.StatusBadRequest
	case errors.Is(err, application.ReceiverNotFound):
		code = http.StatusNotFound
	case errors.Is(err, application.CannotSentCoinsToYourself):
		code = http.StatusBadRequest
	case errors.Is(err, application.MerchNotFound):
		code = http.StatusNotFound
	}

	resp = ErrorResponse{Errors: err.Err.Error()}

	return
}

func (e *ErrorsMiddleware) httpErrorFromDomainError(err domain.Error) (resp ErrorResponse, code int) {
	resp = ErrorResponse{Errors: err.Err.Error()}
	code = http.StatusBadRequest
	return
}

func (e *ErrorsMiddleware) ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	resp := ErrorResponse{Errors: http.StatusText(code)}

	switch v := err.(type) {
	case application.Error:
		resp, code = e.httpErrorFromApplicationError(v)
	case domain.Error:
		resp, code = e.httpErrorFromDomainError(v)
	default:
		c.Logger().Error("unexpected error: ", err)
	}

	if err = c.JSON(code, resp); err != nil {
		c.Logger().Error("failed to write error response", err)
	}
}
