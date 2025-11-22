package todo

import "context"

// ListTodosQuery asks for all todos.
type ListTodosQuery struct{}

// ListTodosHandler handles listing todos.
type ListTodosHandler struct {
	repo Repository
}

func NewListTodosHandler(r Repository) *ListTodosHandler {
	return &ListTodosHandler{repo: r}
}

func (h *ListTodosHandler) Handle(ctx context.Context, q ListTodosQuery) ([]Todo, error) {
	return h.repo.List(ctx)
}
