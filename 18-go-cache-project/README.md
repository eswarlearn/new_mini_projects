# Simple Cache Memory Project in Go

A high-performance, thread-safe in-memory cache implementation in Go with support for TTL (Time To Live), LRU eviction, and comprehensive statistics tracking.

## Features

✅ **Thread-Safe Operations** - Uses `sync.RWMutex` for concurrent access  
✅ **TTL Support** - Automatic expiration of cache items  
✅ **LRU Eviction** - Automatically removes oldest items when cache is full  
✅ **Cache Statistics** - Track hits, misses, evictions, and hit rate  
✅ **Background Cleanup** - Periodic removal of expired items  
✅ **Flexible Configuration** - Configurable max size and default TTL  
✅ **Never-Expiring Items** - Support for permanent cache entries (TTL = 0)

## Project Structure

```
18-go-cache-project/
├── go.mod          # Go module file
├── main.go         # Cache implementation with examples
└── README.md       # This file
```

## Cache Operations

### Core Operations

- **Set(key, value)** - Add/update item with default TTL
- **SetWithTTL(key, value, ttl)** - Add/update item with custom TTL
- **Get(key)** - Retrieve item from cache
- **Delete(key)** - Remove specific item
- **Clear()** - Remove all items from cache

### Utility Functions

- **Size()** - Get current number of items
- **Keys()** - List all cache keys
- **GetStats()** - Get cache statistics
- **Display()** - Print all cache contents

## Usage Examples

### Basic Usage

```go
// Create a cache with max 100 items and 5 minute default TTL
cache := NewCache(100, 5*time.Minute)

// Set values
cache.Set("user:123", "John Doe")
cache.Set("session:abc", tokenData)

// Get values
if value, found := cache.Get("user:123"); found {
    fmt.Println("Found:", value)
}
```

### Custom TTL

```go
// Set item with 10 second expiration
cache.SetWithTTL("temp:data", "expires soon", 10*time.Second)

// Set item that never expires
cache.SetWithTTL("config:db", "connection_string", 0)
```

### Statistics Tracking

```go
stats := cache.GetStats()
fmt.Printf("Hit Rate: %s\n", stats["hitRate"])
fmt.Printf("Total Hits: %d\n", stats["hits"])
fmt.Printf("Total Misses: %d\n", stats["misses"])
fmt.Printf("Evictions: %d\n", stats["evictions"])
```

## How It Works

### Thread Safety
- Uses `sync.RWMutex` to allow multiple concurrent readers or single writer
- Read operations (Get, Size, Keys, etc.) use `RLock()`
- Write operations (Set, Delete, Clear) use `Lock()`

### TTL & Expiration
- Items store expiration timestamp in nanoseconds
- Background goroutine runs every minute to cleanup expired items
- Get operations check expiration before returning values
- TTL of 0 means item never expires

### LRU Eviction
- When cache reaches max size, oldest item is evicted
- Eviction prioritizes items with earliest expiration time
- Eviction counter tracks total evictions for statistics

### Cache Statistics
- **Hits**: Successful Get operations
- **Misses**: Failed Get operations (not found or expired)
- **Evictions**: Items removed due to size limit
- **Hit Rate**: Percentage of successful lookups

## Running the Project

```bash
# Navigate to project directory
cd 18-go-cache-project

# Run the program
go run main.go

# Build executable
go build -o cache main.go

# Run executable
./cache
```

## Example Output

```
🚀 Simple Cache Memory Project in Go

--- Example 1: Basic Operations ---
Found: Alice
user:999 not found (cache miss)

=== Cache Contents ===
  user:1: Alice (expires in 30s)
  user:2: Bob (expires in 30s)
  product:101: Laptop (expires in 30s)
=====================

--- Final Statistics ---
  hits: 15
  misses: 8
  evictions: 3
  hitRate: 65.22%
  size: 5
  maxSize: 5

✅ Cache demonstration complete!
```

## Use Cases

This cache implementation is ideal for:

- **API Response Caching** - Cache expensive API calls
- **Session Management** - Store user session data with TTL
- **Database Query Caching** - Reduce database load
- **Rate Limiting** - Track API request counts per user
- **Configuration Storage** - Store app config with hot reload
- **Computation Results** - Cache expensive calculation results

## Performance Considerations

- **Memory Usage**: Each item stores key, value, and expiration timestamp
- **Concurrency**: Read-heavy workloads benefit from RWMutex
- **Cleanup**: Background goroutine runs every 1 minute
- **Eviction**: O(n) complexity for finding oldest item (can be optimized with LRU linked list)

## Future Enhancements

Potential improvements:
- True LRU implementation with doubly-linked list for O(1) eviction
- Persistence to disk for cache recovery
- Metrics export (Prometheus integration)
- Configurable cleanup interval
- Support for cache warming
- Max memory limit instead of item count
- Batch operations (GetMultiple, SetMultiple)

## License

Free to use for educational and commercial projects.

## Author

Created as a mini project to demonstrate cache implementation in Go.
