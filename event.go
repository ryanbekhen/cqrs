package cqrs

import (
	"context"
	"reflect"
	"sync"
)

var (
	muEvents      sync.RWMutex
	eventHandlers = make(map[reflect.Type][]interface{})
)

type Event interface{}

type EventHandler[E Event] interface {
	Handle(ctx context.Context, e E) error
}

func RegisterEvent[E Event](h EventHandler[E]) {
	t := reflect.TypeOf(*new(E))
	muEvents.Lock()
	defer muEvents.Unlock()
	eventHandlers[t] = append(eventHandlers[t], h)
}

func Publish[E Event](ctx context.Context, e E) error {
	t := reflect.TypeOf(e)
	muEvents.RLock()
	defer muEvents.RUnlock()
	handlers := eventHandlers[t]
	for _, h := range handlers {
		if err := h.(EventHandler[E]).Handle(ctx, e); err != nil {
			return err
		}
	}
	return nil
}
