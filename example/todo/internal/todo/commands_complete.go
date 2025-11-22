package todo

import (
	"context"
	"time"
)

// CompleteTodoCommand marks a todo as completed.
type CompleteTodoCommand struct {
	ID string
}

// CompleteTodoHandler handles completing todos.
type CompleteTodoHandler struct {
	repo Repository
}

func NewCompleteTodoHandler(r Repository) *CompleteTodoHandler {
	return &CompleteTodoHandler{repo: r}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, cmd CompleteTodoCommand) (Todo, error) {
	// load
	t, err := h.repo.GetByID(ctx, cmd.ID)
	if err != nil {
		return Todo{}, err
	}
	if t.Completed {
		return t, nil // idempotent
	}
	now := time.Now()
	t.Completed = true
	t.CompletedAt = &now
	return h.repo.Update(ctx, t)
}
