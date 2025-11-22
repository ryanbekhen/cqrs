package cqrs

import (
	"context"
	"testing"
)

type TestQuery struct{ Val int }
type TestQueryHandler struct{}

func (h *TestQueryHandler) Handle(ctx context.Context, q TestQuery) (int, error) {
	return q.Val * 2, nil
}

func TestQueryDispatch(t *testing.T) {
	handler := &TestQueryHandler{}
	RegisterQuery[TestQuery, int](handler)

	res, err := DispatchQuery[TestQuery, int](context.Background(), TestQuery{Val: 3})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res != 6 {
		t.Fatalf("expected 6, got %v", res)
	}
}
