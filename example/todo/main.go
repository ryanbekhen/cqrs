package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanbekhen/cqrs/example/todo/api"
	"github.com/ryanbekhen/cqrs/example/todo/internal/todo"
)

func main() {
	repo := todo.NewMemoryRepo()

	handlers := &api.Handlers{
		Create:   todo.NewCreateTodoHandler(repo),
		Get:      todo.NewGetTodoHandler(repo),
		List:     todo.NewListTodosHandler(repo),
		Complete: todo.NewCompleteTodoHandler(repo),
	}

	r := api.NewRouter(handlers)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("todo service listening at :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	fmt.Println("server stopped")
}
