package cqrs

import (
	"context"
	"testing"
)

type BenchmarkCommand struct{ Val int }
type BenchmarkCommandHandler struct{}

func (h *BenchmarkCommandHandler) Handle(ctx context.Context, cmd BenchmarkCommand) error {
	return nil
}

func BenchmarkCommandDispatch(b *testing.B) {
	handler := &BenchmarkCommandHandler{}
	RegisterCommand[BenchmarkCommand](handler)
	ctx := context.Background()
	cmd := BenchmarkCommand{Val: 42}

	b.ResetTimer()
	b.ReportAllocs() // <-- laporan alokasi memori

	for i := 0; i < b.N; i++ {
		_ = Dispatch(ctx, cmd)
	}
}
