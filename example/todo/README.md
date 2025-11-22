Vertical-slice Todo example

This example demonstrates a small Todo service structured as vertical slices. It uses an in-memory repository and stdlib HTTP server.

Run:

```bash
cd example/todo
go run .
```

Test:

```bash
cd example/todo
go test ./...
```

API:
- POST /todos            -> create
- GET  /todos            -> list
- GET  /todos/{id}       -> get
- POST /todos/{id}/complete -> mark completed

