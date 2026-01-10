# Goroutines Exercise: Parallel Ad Fetching

This exercise demonstrates how to use **goroutines** to fetch data from multiple downstream services in parallel, improving performance compared to sequential requests.

## Overview

The exercise simulates fetching advertisements from three different ad sources (Google, Amazon, and Meta) concurrently using Go's goroutines.

## Key Components

### Ad Source (`ads.go`)
- Defines `AdSource` type representing different ad providers (Google, Meta, Amazon)
- Implements the `Stringer` interface for human-readable source names
- Defines the `Ad` struct to represent individual advertisements

### Parallel Fetcher (`adsfetch.go`)
- **`fetchAds()`**: Simulates an HTTP request to an ad source with random delays (100-200ms)
- **`FetchAdsFromMultipleSources()`**: Fetches ads from all sources concurrently using goroutines

## Concurrency Patterns Used

1. **`sync.WaitGroup`**: Coordinates waiting for all goroutines to complete
   - `wg.Add(1)` before launching each goroutine
   - `wg.Done()` when each goroutine finishes
   - `wg.Wait()` blocks until all goroutines complete

2. **`sync.Mutex`**: Protects shared map access
   - `mu.Lock()` before writing to the result map
   - `mu.Unlock()` after writing completes
   - Prevents race conditions when multiple goroutines write simultaneously

3. **Goroutine with closure**: Each ad source is passed as a parameter to avoid closure variable issues

## Benefits

Running requests in parallel reduces total execution time from ~600ms (3 sources × ~200ms each) to ~200ms (the longest individual request).

## Running the Test

```bash
go test -v ./goroutines
```
