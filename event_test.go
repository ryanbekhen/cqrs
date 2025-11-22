package cqrs

import (
	"context"
	"testing"
)

type TestEvent struct{ Val int }
type TestEventHandler struct{}

func (h *TestEventHandler) Handle(ctx context.Context, e TestEvent) error {
	return nil
}

func TestEventPublish(t *testing.T) {
	handler := &TestEventHandler{}
	RegisterEvent[TestEvent](handler)

	err := Publish[TestEvent](context.Background(), TestEvent{Val: 5})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
