package todo

import (
	"context"
	"testing"
)

func TestMemoryRepo_CreateGetListUpdate(t *testing.T) {
	r := NewMemoryRepo()
	ctx := context.Background()
	// create
	todoIn := Todo{Title: "test", Description: "desc"}
	created, err := r.Create(ctx, todoIn)
	if err != nil {
		t.Fatalf("create error: %v", err)
	}
	if created.ID == "" {
		t.Fatalf("expected id assigned")
	}
	// get
	tGot, err := r.GetByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("get error: %v", err)
	}
	if tGot.Title != "test" {
		t.Fatalf("unexpected title: %s", tGot.Title)
	}
	// list
	list, err := r.List(ctx)
	if err != nil {
		t.Fatalf("list error: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 item, got %d", len(list))
	}
	// update (complete)
	now := list[0]
	now.Completed = true
	updated, err := r.Update(ctx, now)
	if err != nil {
		t.Fatalf("update error: %v", err)
	}
	if !updated.Completed {
		t.Fatalf("expected completed true")
	}
}
