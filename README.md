# Text & Password Analyzer (Checkpoint 1)

## 1. Project Overview

This project is an in-memory Text & Password Analyzer built for Checkpoint 1. It calculates basic metrics for text strings and passwords without any external dependencies, databases, or HTTP servers.

### Main Components

- **`model.go`**: Contains the `Task` and `TextResult` structures, status constants (`PENDING`, `COMPLETED`), and domain errors.
- **`repository.go`**: `TaskRepository` interface and `InMemoryRepository` implementation using a `map[string]*Task`.
- **`processor.go`**: `TaskProcessor` interface and `TextProcessor` implementation that validates input, calculates metrics, and updates task status.
- **`main.go`**: Demonstration of creating, saving, and processing tasks.
- **`task_test.go`**: Automated unit tests (including table-driven tests) verifying processing, error handling, and repository methods.

---

## 2. Task Lifecycle

Tasks follow a simple two-state lifecycle:

```
[PENDING]  ─── Process() ───►  [COMPLETED]
```

- When created, a task is in the `PENDING` state.
- Upon calling `Process()`, input is validated:
  - If input is empty or contains only whitespace, `ErrEmptyInput` is returned.
  - If the task is already `COMPLETED`, `ErrAlreadyDone` is returned.
- Upon successful processing, word and character counts are calculated, `IsSecure` is set (`true` if length $\ge 8$), and the task transitions to `COMPLETED`.

---

## 3. Commands

Run all commands from the project directory:

```bash
cd ~/Downloads/KBTU2025-2026/go_project
```

### Check code formatting:

```bash
gofmt -l .
```

### Run static analyzer (linter):

```bash
go vet ./...
```

### Run unit tests:

```bash
go test -v ./...
```

### Run the program:

```bash
go run .
```
