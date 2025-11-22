package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryanbekhen/cqrs/example/todo/internal/todo"
)

func TestHandlers_CRUD(t *testing.T) {
	repo := todo.NewMemoryRepo()
	h := &Handlers{
		Create:   todo.NewCreateTodoHandler(repo),
		Get:      todo.NewGetTodoHandler(repo),
		List:     todo.NewListTodosHandler(repo),
		Complete: todo.NewCompleteTodoHandler(repo),
	}
	r := NewRouter(h)

	// create
	reqBody := map[string]string{"title": "buy milk", "description": "2L"}
	b, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(b))
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", res.StatusCode)
	}
	var created todo.Todo
	_ = json.NewDecoder(res.Body).Decode(&created)
	if created.ID == "" {
		t.Fatalf("expected id")
	}
	// get
	req = httptest.NewRequest(http.MethodGet, "/todos/"+created.ID, nil)
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	res = rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 get, got %d", res.StatusCode)
	}
	var got todo.Todo
	_ = json.NewDecoder(res.Body).Decode(&got)
	if got.ID != created.ID {
		t.Fatalf("id mismatch")
	}
	// list
	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	res = rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 list, got %d", res.StatusCode)
	}
	var list []todo.Todo
	_ = json.NewDecoder(res.Body).Decode(&list)
	if len(list) != 1 {
		t.Fatalf("expected list len 1")
	}
	// complete
	req = httptest.NewRequest(http.MethodPost, "/todos/"+created.ID+"/complete", nil)
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	res = rec.Result()
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		t.Fatalf("expected 200 complete, got %d - %s", res.StatusCode, string(body))
	}
	var done todo.Todo
	_ = json.NewDecoder(res.Body).Decode(&done)
	if !done.Completed {
		t.Fatalf("expected completed true")
	}
	if done.CompletedAt == nil {
		t.Fatalf("expected CompletedAt set")
	}
	_ = res.Body.Close()
	_ = res.Body.Close()
}
