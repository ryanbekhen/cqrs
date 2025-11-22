package todo

import "context"

// CreateTodoCommand is the input for creating a todo.
type CreateTodoCommand struct {
	Title       string
	Description string
}

// CreateTodoHandler handles creating todos.
type CreateTodoHandler struct {
	repo Repository
}

func NewCreateTodoHandler(r Repository) *CreateTodoHandler {
	return &CreateTodoHandler{repo: r}
}

func (h *CreateTodoHandler) Handle(ctx context.Context, cmd CreateTodoCommand) (Todo, error) {
	if cmd.Title == "" {
		return Todo{}, errorString("title required")
	}
	in := Todo{
		Title:       cmd.Title,
		Description: cmd.Description,
	}
	return h.repo.Create(ctx, in)
}
