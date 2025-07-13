# Go Concurrency

Learning Go concurrency concepts through hands-on practice, following along with a structured course. This repository contains my notes and code as I work through goroutines, wait groups, mutexes, and channels.

## Table of Contents

- [Go Concurrency](#go-concurrency)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Clone and Run](#clone-and-run)
  - [Concepts Covered](#concepts-covered)
    - [Goroutines](#goroutines)
    - [Wait Groups](#wait-groups)
    - [Mutexes](#mutexes)
      - [Race Conditions](#race-conditions)
    - [Channels](#channels)
  - [Running the Code](#running-the-code)
  - [Testing for Race Conditions](#testing-for-race-conditions)
  - [Project Structure](#project-structure)
  - [Notes](#notes)

## About

This repository documents my journey learning Go's concurrency primitives. Each concept includes practical code and notes from working through real examples.

## Getting Started

### Prerequisites

- Go 1.19 or higher

### Clone and Run

```bash
git clone https://github.com/kmulqueen/go-concurrency.git
cd go-concurrency
```

## Concepts Covered

### Goroutines

Lightweight threads managed by the Go runtime. The foundation of Go's concurrency model.

### Wait Groups

Synchronization primitive for waiting on multiple goroutines to complete. Part of the `sync` package.

**Key Methods:**

- `Add()` - increment the counter
- `Done()` - decrement the counter (use with `defer`)
- `Wait()` - block until counter reaches zero

### Mutexes

**Mutex = "Mutual Exclusion"** - allows us to deal with race conditions. Relatively simple to use when dealing with shared resources and concurrent/parallel goroutines. Uses `Lock()` and `Unlock()` methods. We can test for race conditions when running code or testing it.

**Key Concepts:**

- Protects shared state from concurrent access
- Always use `defer` to unlock
- Prevents data races in concurrent code

#### Race Conditions

- Race conditions occur when multiple goroutines try to access the same data.
- Can be difficult to spot when reading code.
- Go allows us to check for them when running a program, or when testing our code with `go test` (see [Testing for Race Conditions](#testing-for-race-conditions)).

### Channels

- Channels are a means of having goroutines share data.
- They can talk to each other.
  - Can be uni-directional (send-only or receive-only) or bi-directional.
- This is Go's philosophy of having things share memory by communicating, rather than communicating by sharing memory.
- The Producer/Consumer problem.

## Running the Code

```bash
# Run wait groups example
cd wait-groups
go run main.go

# Run wait groups challenge
cd wait-groups/challenge-1
go run main.go
```

## Testing for Race Conditions

Go provides built-in race detection. Use the `-race` flag:

```bash
# Run with race detection
go run -race main.go

# Test with race detection
go test -race

# Test specific package
go test -race ./wait-groups
```

## Project Structure

```
go-concurrency/
├── README.md
└── wait-groups/
    ├── main.go          # Basic wait groups example
    ├── main_test.go     # Tests for wait groups
    ├── go.mod
    └── challenge-1/
        ├── main.go      # Wait groups challenge
        ├── main_test.go # Challenge tests
        └── go.mod
```

## Notes

This is a learning repository following a course on Go concurrency. Currently focusing on the foundational concepts before moving on to more advanced topics like channels and select statements.
