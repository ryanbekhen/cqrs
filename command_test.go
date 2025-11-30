package cqrs

import (
	"context"
	"testing"
)

type TestCommand struct{ Val int }

type TestCommandHandler struct{}

func (h *TestCommandHandler) Handle(ctx context.Context, cmd TestCommand) (any, error) {
	return nil, nil
}

func TestCommandDispatch(t *testing.T) {
	handler := &TestCommandHandler{}
	RegisterCommand[TestCommand](handler)

	_, err := DispatchCommand[TestCommand, any](context.Background(), TestCommand{Val: 1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
