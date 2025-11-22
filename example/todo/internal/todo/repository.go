package todo

import "context"

// Repository defines storage operations for todos.
type Repository interface {
	Create(ctx context.Context, t Todo) (Todo, error)
	GetByID(ctx context.Context, id string) (Todo, error)
	List(ctx context.Context) ([]Todo, error)
	Update(ctx context.Context, t Todo) (Todo, error)
}
