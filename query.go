package cqrs

import (
	"context"
	"reflect"
	"sync"
)

var (
	muQueries     sync.RWMutex
	queryHandlers = make(map[reflect.Type]interface{})
)

type Query interface{}

type QueryHandler[Q Query, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}

func RegisterQuery[Q Query, R any](h QueryHandler[Q, R]) {
	t := reflect.TypeOf(*new(Q))
	muQueries.Lock()
	defer muQueries.Unlock()
	queryHandlers[t] = h
}

func DispatchQuery[Q Query, R any](ctx context.Context, q Q) (R, error) {
	var zero R
	t := reflect.TypeOf(q)
	muQueries.RLock()
	defer muQueries.RUnlock()
	h, ok := queryHandlers[t]
	if !ok {
		return zero, ErrQueryHandlerNotFound
	}
	return h.(QueryHandler[Q, R]).Handle(ctx, q)
}
