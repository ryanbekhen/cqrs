package todo

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// MemoryRepo is an in-memory implementation of Repository.
type MemoryRepo struct {
	mu   sync.RWMutex
	seq  int64
	data map[string]Todo
}

// NewMemoryRepo creates a new in-memory repository.
func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		data: make(map[string]Todo),
	}
}

func (r *MemoryRepo) Create(ctx context.Context, t Todo) (Todo, error) {
	id := atomic.AddInt64(&r.seq, 1)
	t.ID = strconv.FormatInt(id, 10)
	t.CreatedAt = time.Now()

	r.mu.Lock()
	defer r.mu.Unlock()
	// copy to avoid mutation from caller
	r.data[t.ID] = t
	return t, nil
}

func (r *MemoryRepo) GetByID(ctx context.Context, id string) (Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if t, ok := r.data[id]; ok {
		return t, nil
	}
	return Todo{}, ErrNotFound
}

func (r *MemoryRepo) List(ctx context.Context) ([]Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Todo, 0, len(r.data))
	for _, v := range r.data {
		out = append(out, v)
	}
	return out, nil
}

func (r *MemoryRepo) Update(ctx context.Context, t Todo) (Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[t.ID]; !ok {
		return Todo{}, ErrNotFound
	}
	r.data[t.ID] = t
	return t, nil
}
