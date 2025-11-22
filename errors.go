package cqrs

import "errors"

var (
	ErrCommandHandlerNotFound = errors.New("command handler not found")
	ErrQueryHandlerNotFound   = errors.New("query handler not found")
)
