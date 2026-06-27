# kev

kev is a lightweight in-memory key-value database written in Go.

The project is built for learning low-level data structures, concurrency, networking, and database internals. It implements a simple TCP server with an in-memory storage engine backed by an LRU eviction policy and key expiration.

## Features

- In-memory key-value storage
- LRU eviction
- Key expiration (TTL)
- Concurrent access using `sync.RWMutex`
- TCP server
- Interactive CLI
- Generic doubly linked list implementation

## Project Structure

```
.
├── command.go
├── list/
├── server/
├── store.go
├── main.go
└── ...
```

## Running

```bash
go run .
```

## Running Tests

```bash
go test -v
```

Race detector:

```bash
go test -race
```

## Example Commands

```text
SET name kev
GET name
DELETE name
EXIT
```

## Current Architecture

```
          TCP Client
               │
               ▼
         Command Parser
               │
               ▼
            Store
         ┌──────────┐
         │   Map    │
         └──────────┘
               │
               ▼
        Generic LRU List
```

## Roadmap

### Core

- [x] In-memory key-value storage
- [x] TCP server
- [x] CLI
- [x] Generic doubly linked list
- [x] LRU eviction
- [x] TTL expiration
- [x] Unit tests

### Data Types

- [ ] Integer values
- [ ] Lists
- [ ] Hashes
- [ ] Sets
- [ ] Sorted sets

### Commands

- [ ] EXISTS
- [ ] KEYS
- [ ] CLEAR
- [ ] EXPIRE
- [ ] TTL
- [ ] RENAME
- [ ] INCR / DECR

### Networking

- [ ] Multiple concurrent clients
- [ ] Connection pool
- [ ] RESP protocol compatibility
- [ ] Pipelining

### Performance

- [ ] Benchmarks
- [ ] Memory profiling
- [ ] CPU profiling
- [ ] Sharded locking
- [ ] Efficient expiration strategy (min-heap or timing wheel)

### Reliability

- [ ] Better error handling
- [ ] Configuration file
- [ ] Graceful shutdown
- [ ] Logging
- [ ] Metrics endpoint

### Developer Experience

- [ ] Integration tests
- [ ] Benchmark suite
- [ ] GitHub Actions
- [ ] Docker support
- [ ] API documentation

## Goals

The goal of kev is not to become a production database, but to explore the internal design of in-memory databases and practice systems programming concepts in Go.
