package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ryanbekhen/cqrs"
)

type PingCommand struct{ Msg string }
type PongEvent struct{ Reply string }

type PingHandler struct{}

func (h *PingHandler) Handle(ctx context.Context, cmd *PingCommand) (any, error) {
	// simulate work
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Handled Ping:", cmd.Msg)
	if err := cqrs.Publish(ctx, PongEvent{Reply: "pong: " + cmd.Msg}); err != nil {
		return nil, err
	}
	return nil, nil
}

type PongHandler struct{}

func (h *PongHandler) Handle(ctx context.Context, e PongEvent) error {
	fmt.Println("Received Pong:", e.Reply)
	return nil
}

func main() {
	ctx := context.Background()

	// register event handler
	cqrs.RegisterEvent(&PongHandler{})
	// register command handler
	cqrs.RegisterCommand(&PingHandler{})

	const numCommands = 10
	var wg sync.WaitGroup
	for i := 1; i <= numCommands; i++ {
		i := i // capture loop variable
		wg.Go(func() {
			cmd := PingCommand{Msg: fmt.Sprintf("msg-%d", i)}
			if _, err := cqrs.DispatchCommand[*PingCommand, any](ctx, &cmd); err != nil {
				fmt.Println("Error dispatching command:", err)
			}
		})
	}

	wg.Wait()
	fmt.Println("All commands processed")
}
