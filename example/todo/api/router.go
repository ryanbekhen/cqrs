package api

import (
	"net/http"
)

// NewRouter wires HTTP routes for the todo API.
func NewRouter(h *Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateTodo(w, r)
		case http.MethodGet:
			h.ListTodos(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// patterns for /todos/{id} and /todos/{id}/complete
	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetTodo(w, r)
			return
		}
		// POST /todos/{id}/complete
		if r.Method == http.MethodPost && len(r.URL.Path) > len("/todos/") && r.URL.Path[len(r.URL.Path)-len("/complete"):] == "/complete" {
			h.CompleteTodo(w, r)
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
	})
	return mux
}
