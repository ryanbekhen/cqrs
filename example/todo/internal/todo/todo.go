package todo

import "time"

// Todo is the domain model for a task.
type Todo struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// Errors
var (
	ErrNotFound = errorString("todo: not found")
)

type errorString string

func (e errorString) Error() string { return string(e) }
