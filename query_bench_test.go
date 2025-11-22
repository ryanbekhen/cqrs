package cqrs

import (
	"context"
	"testing"
)

type BenchmarkQuery struct{ Val int }
type BenchmarkQueryHandler struct{}

func (h *BenchmarkQueryHandler) Handle(ctx context.Context, q BenchmarkQuery) (int, error) {
	return q.Val * 2, nil
}

func BenchmarkQueryDispatch(b *testing.B) {
	handler := &BenchmarkQueryHandler{}
	RegisterQuery[BenchmarkQuery, int](handler)
	ctx := context.Background()
	q := BenchmarkQuery{Val: 10}

	b.ResetTimer()
	b.ReportAllocs() // <-- laporan alokasi memori

	for i := 0; i < b.N; i++ {
		_, _ = DispatchQuery[BenchmarkQuery, int](ctx, q)
	}
}
