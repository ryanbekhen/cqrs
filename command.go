package cqrs

import (
	"context"
	"reflect"
	"sync"
)

var (
	muCommands  sync.RWMutex
	cmdHandlers = make(map[reflect.Type]interface{})
)

type Command interface{}

type CommandHandler[C Command, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

func RegisterCommand[C Command, R any](h CommandHandler[C, R]) {
	var c C
	t := reflect.TypeOf(c)
	muCommands.Lock()
	defer muCommands.Unlock()
	cmdHandlers[t] = h
}

func DispatchCommand[C Command, R any](ctx context.Context, cmd C) (R, error) {
	var zero R

	t := reflect.TypeOf(cmd)
	muCommands.RLock()
	h, ok := cmdHandlers[t]
	muCommands.RUnlock()
	if !ok {
		return zero, ErrCommandHandlerNotFound
	}
	return h.(CommandHandler[C, R]).Handle(ctx, cmd)
}
