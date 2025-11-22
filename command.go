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

type CommandHandler[C Command] interface {
	Handle(ctx context.Context, cmd C) error
}

func RegisterCommand[C Command](h CommandHandler[C]) {
	t := reflect.TypeOf(*new(C))
	muCommands.Lock()
	defer muCommands.Unlock()
	cmdHandlers[t] = h
}

func Dispatch[C Command](ctx context.Context, cmd C) error {
	t := reflect.TypeOf(cmd)
	muCommands.RLock()
	defer muCommands.RUnlock()
	h, ok := cmdHandlers[t]
	if !ok {
		return ErrCommandHandlerNotFound
	}
	return h.(CommandHandler[C]).Handle(ctx, cmd)
}
