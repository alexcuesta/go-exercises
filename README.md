# Go Exercises

A collection of Go programming exercises focusing on concurrency patterns and data structures.

## Exercises

### 1. Goroutines ([goroutines/](goroutines/))
Learn parallel processing by fetching ads from multiple sources concurrently.

**Concepts covered:**
- Goroutines
- `sync.WaitGroup` for coordinating concurrent operations
- `sync.Mutex` for protecting shared resources
- Parallel API requests simulation

See [goroutines/README.md](goroutines/README.md) for details.

### 2. Channels ([channels/](channels/))
Implement a worker pool pattern using Go channels.

**Concepts covered:**
- Buffered and unbuffered channels
- Channel direction (`<-chan`, `chan<-`)
- Worker pool pattern
- Proper channel closing to avoid deadlocks

See [channels/EXERCISE.md](channels/EXERCISE.md) for the exercise description.

### 3. Graph ([graph/](graph/))
Implement a graph data structure with Breadth-First Search (BFS) traversal.

**Concepts covered:**
- Adjacency list representation
- BFS algorithm implementation
- Error handling
- Custom data structures

## Running the Exercises

Run all tests:
```bash
go test ./...
```

Run a specific exercise:
```bash
go test -v ./goroutines
go test -v ./channels
go test -v ./graph
```

## Requirements

- Go 1.25.3 or later