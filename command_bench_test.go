package cqrs

import (
	"context"
	"testing"
)

type BenchmarkCommand struct{ Val int }
type BenchmarkCommandHandler struct{}

func (h *BenchmarkCommandHandler) Handle(ctx context.Context, cmd BenchmarkCommand) (interface{}, error) {
	return nil, nil
}

func BenchmarkCommandDispatch(b *testing.B) {
	handler := &BenchmarkCommandHandler{}
	RegisterCommand[BenchmarkCommand](handler)
	ctx := context.Background()
	cmd := BenchmarkCommand{Val: 42}

	b.ResetTimer()
	b.ReportAllocs() // <-- laporan alokasi memori

	for i := 0; i < b.N; i++ {
		_, _ = DispatchCommand[BenchmarkCommand, any](ctx, cmd)
	}
}
