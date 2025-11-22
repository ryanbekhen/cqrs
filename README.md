# go-cqrs

[English](./README.md) | [Bahasa Indonesia](./README-id.md)

Minimal implementation of the CQRS (Command Query Responsibility Segregation) pattern for Go.

## Overview

`go-cqrs` is a lightweight library for applying the CQRS pattern in Go applications. It provides a simple API to register and dispatch commands, register and publish events, and register and dispatch queries.

## Features

- Register and Dispatch Commands (one handler per command type)
- Register and Publish Events (multiple handlers per event type)
- Register and Dispatch Queries (one handler per query, returns a result)
- Generics-based API for working with concrete types
- Safe for concurrent access (uses mutexes)

## Installation

Ensure Go (>= 1.18) is installed. Add the module to your project:

```bash
go get github.com/ryanbekhen/cqrs
```

Or require the module directly in your `go.mod`:

```go
require github.com/ryanbekhen/cqrs latest
```

## Usage Example

```go
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

// Command handler for PingCommand
type PingHandler struct{}
func (h *PingHandler) Handle(ctx context.Context, cmd PingCommand) error {
    // simulate work
    time.Sleep(50 * time.Millisecond)
    fmt.Println("Handled Ping:", cmd.Msg)
    return cqrs.Publish(ctx, PongEvent{Reply: "pong: " + cmd.Msg})
}

// Event handler for PongEvent
type PongHandler struct{}
func (h *PongHandler) Handle(ctx context.Context, e PongEvent) error {
    fmt.Println("Received Pong:", e.Reply)
    return nil
}

func main() {
    ctx := context.Background()

    // register handlers
    cqrs.RegisterEvent(&PongHandler{})
    cqrs.RegisterCommand(&PingHandler{})

    const numCommands = 10
    var wg sync.WaitGroup
    wg.Add(numCommands)

    for i := 1; i <= numCommands; i++ {
        i := i
        go func() {
            defer wg.Done()
            cmd := PingCommand{Msg: fmt.Sprintf("msg-%d", i)}
            if err := cqrs.Dispatch(ctx, cmd); err != nil {
                fmt.Println("Error dispatching command:", err)
            }
        }()
    }

    wg.Wait()
    fmt.Println("All commands processed")
}
```

## API (Summary)

Core types:

- `type Command interface{}`
- `type Event interface{}`
- `type Query interface{}`

Commands:
- `RegisterCommand[C Command](h CommandHandler[C])` — register a handler for command type C.
- `Dispatch[C Command](ctx context.Context, cmd C) error` — dispatch a command to its handler.

Events:
- `RegisterEvent[E Event](h EventHandler[E])` — register an event handler (multiple handlers allowed).
- `Publish[E Event](ctx context.Context, e E) error` — publish an event to all registered handlers.

Queries:
- `RegisterQuery[Q Query, R any](h QueryHandler[Q, R])` — register a handler for query Q that returns R.
- `DispatchQuery[Q Query, R any](ctx context.Context, q Q) (R, error)` — execute the query and receive the result.

## Common Errors

- `ErrCommandHandlerNotFound` — no handler registered for the dispatched command.
- `ErrQueryHandlerNotFound` — no handler registered for the dispatched query.

## Implementation Notes

- Type-to-handler mapping is implemented with `reflect`.
- `sync.RWMutex` is used to protect handler maps for concurrent access.
- The design is intentionally minimal to stay readable and easy to extend (e.g., add middleware, tracing, or async dispatch).

## Testing

Run tests with:

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a new branch: `git checkout -b feature-name`
3. Commit your changes
4. Open a Pull Request and describe your changes

## License

This project is licensed under the MIT License — see the `LICENSE` file.

