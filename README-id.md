# go-cqrs

[![Go Reference](https://pkg.go.dev/badge/github.com/ryanbekhen/cqrs.svg)](https://pkg.go.dev/github.com/ryanbekhen/cqrs)
[![Coverage Status](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/ryanbekhen/cqrs/coverage-badge/.badges/coverage.json)](https://codecov.io/gh/ryanbekhen/cqrs)
[![License](https://img.shields.io/github/license/ryanbekhen/cqrs?style=flat-square)](LICENSE)

[English](./README.md) | [Bahasa Indonesia](./README-id.md)

Implementasi minimal pola CQRS (Command Query Responsibility Segregation) untuk Go.

## Ringkasan

`go-cqrs` adalah pustaka ringan untuk menerapkan pola CQRS pada aplikasi Go. Pustaka ini menyediakan API sederhana untuk mendaftarkan dan men-dispatch command, mendaftarkan dan mem-publish event, serta mendaftarkan dan men-dispatch query.

## Fitur

- Register dan Dispatch Command (satu handler per tipe command)
- Register dan Publish Event (banyak handler per tipe event)
- Register dan Dispatch Query (satu handler per query, mengembalikan hasil)
- API berbasis generics untuk penggunaan tipe konkret
- Aman untuk akses konkuren (menggunakan mutex)

## Instalasi

Pastikan Go (>= 1.18) terpasang. Tambahkan modul ke proyek Anda:

```bash
go get github.com/ryanbekhen/cqrs
```

Atau gunakan modul langsung di `go.mod` Anda:

```go
require github.com/ryanbekhen/cqrs latest
```

## Contoh Penggunaan

Lihat di folder `example`.

## API (Ringkasan)

Tipe dasar:

- `type Command interface{}`
- `type Event interface{}`
- `type Query interface{}`

Command:
- `RegisterCommand[C Command, R any](h CommandHandler[C, R])` — daftarkan handler untuk command bertipe C.
- `DispatchCommand[C Command, R any](ctx context.Context, cmd C) (R, error)` — kirim command ke handler.

Event:
- `RegisterEvent[E Event](h EventHandler[E])` — daftarkan event handler (banyak handler diperbolehkan).
- `Publish[E Event](ctx context.Context, e E) error` — publikasikan event ke semua handler.

Query:
- `RegisterQuery[Q Query, R any](h QueryHandler[Q, R])` — daftarkan handler untuk query Q yang mengembalikan R.
- `DispatchQuery[Q Query, R any](ctx context.Context, q Q) (R, error)` — jalankan query dan terima hasil.

## Error Umum

- `ErrCommandHandlerNotFound` — tidak ada handler terdaftar untuk command yang di-dispatch.
- `ErrQueryHandlerNotFound` — tidak ada handler terdaftar untuk query yang di-dispatch.

## Catatan Implementasi

- Pemetaan tipe ke handler menggunakan `reflect`.
- Menggunakan `sync.RWMutex` untuk akses peta handler secara aman pada kondisi konkuren.
- Desain sengaja minimal agar mudah dibaca dan diperluas (mis. menambahkan middleware, tracing, atau async dispatch jika diperlukan).

## Pengujian

Untuk menjalankan test (jika ada):

```bash
go test ./...
```

## Kontribusi

1. Fork repositori
2. Buat branch baru: `git checkout -b fitur-baru`
3. Commit perubahan Anda
4. Buat Pull Request dan jelaskan perubahan yang diusulkan

## Lisensi

Proyek ini dilisensikan di bawah MIT License — lihat file `LICENSE`.
