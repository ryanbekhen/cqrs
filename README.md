# go-cqrs

[![Go Reference](https://pkg.go.dev/badge/github.com/ryanbekhen/cqrs.svg)](https://pkg.go.dev/github.com/ryanbekhen/cqrs)
[![Coverage Status](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/ryanbekhen/cqrs/coverage-badge/.badges/coverage.json)](https://codecov.io/gh/ryanbekhen/cqrs)
[![License](https://img.shields.io/github/license/ryanbekhen/cqrs?style=flat-square)](LICENSE)

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

See the `example` folder.

## API (Summary)

Core types:

- `type Command interface{}`
- `type Event interface{}`
- `type Query interface{}`

Commands:
- `RegisterCommand[C Command, R any](h CommandHandler[C, R])` — register a handler for command type C.
- `DispatchCommand[C Command, R any](ctx context.Context, cmd C) (R, error)` — dispatch a command to its handler.

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

