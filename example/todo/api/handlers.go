package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ryanbekhen/cqrs/example/todo/internal/todo"
)

// Handlers groups application handlers used by HTTP endpoints.
type Handlers struct {
	Create   *todo.CreateTodoHandler
	Get      *todo.GetTodoHandler
	List     *todo.ListTodosHandler
	Complete *todo.CompleteTodoHandler
}

// request/response payloads
type createReq struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	res, err := h.Create.Handle(r.Context(), todo.CreateTodoCommand{Title: req.Title, Description: req.Description})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *Handlers) ListTodos(w http.ResponseWriter, r *http.Request) {
	res, err := h.List.Handle(r.Context(), todo.ListTodosQuery{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func (h *Handlers) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	// strip trailing elements like /complete
	if idx := strings.Index(id, "/"); idx != -1 {
		id = id[:idx]
	}
	res, err := h.Get.Handle(r.Context(), todo.GetTodoQuery{ID: id})
	if err != nil {
		if err == todo.ErrNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func (h *Handlers) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	// path: /todos/{id}/complete
	p := strings.TrimPrefix(r.URL.Path, "/todos/")
	id := strings.TrimSuffix(p, "/complete")
	id = strings.TrimSuffix(id, "/")
	res, err := h.Complete.Handle(r.Context(), todo.CompleteTodoCommand{ID: id})
	if err != nil {
		if err == todo.ErrNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
