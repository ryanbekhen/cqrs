package cqrs

import (
	"context"
	"testing"
)

type BenchmarkEvent struct{ Val int }
type BenchmarkEventHandler struct{}

func (h *BenchmarkEventHandler) Handle(ctx context.Context, e BenchmarkEvent) error {
	return nil
}

func BenchmarkEventPublish(b *testing.B) {
	handler := &BenchmarkEventHandler{}
	RegisterEvent[BenchmarkEvent](handler)
	ctx := context.Background()
	e := BenchmarkEvent{Val: 100}

	b.ResetTimer()
	b.ReportAllocs() // <-- laporan alokasi memori

	for i := 0; i < b.N; i++ {
		_ = Publish(ctx, e)
	}
}
