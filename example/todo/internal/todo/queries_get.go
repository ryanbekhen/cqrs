package todo

import "context"

// GetTodoQuery asks for a todo by id.
type GetTodoQuery struct {
	ID string
}

// GetTodoHandler handles queries to fetch a todo.
type GetTodoHandler struct {
	repo Repository
}

func NewGetTodoHandler(r Repository) *GetTodoHandler {
	return &GetTodoHandler{repo: r}
}

func (h *GetTodoHandler) Handle(ctx context.Context, q GetTodoQuery) (Todo, error) {
	return h.repo.GetByID(ctx, q.ID)
}
