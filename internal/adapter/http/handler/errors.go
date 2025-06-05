package handler

import "errors"

type Error error

var (
	Unauthorized Error = errors.New("unauthorized")
)
