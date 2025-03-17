# Cache Package
## Overview
The cache package provides a simple, thread-safe cache implementation in Go. It uses a map to store key-value pairs and a mutex for synchronization, ensuring safe access in concurrent environments.

## Features
- **Singleton Pattern**: The cache is implemented as a singleton, ensuring only one instance is created throughout the application lifecycle.
- **Thread Safety**: Uses `sync.RWMutex` to allow multiple readers or a single writer to access the cache simultaneously.
- **Generic Value Type**: Supports storing values of any type (`any`).
- **Error Handling**: Defines a custom error type for type conversion issues.

## Usage
### Getting the Cache Instance
To use the cache, first obtain an instance using the `GetInstance()` function.
```go
cache := cache.GetInstance()
```

### Setting Values
Set a value in the cache using the `Set()` method.
```go
cache.Set("key", "value")
```


### Retrieving Values
Get a value from the cache using the `Get()` method. It returns the value and a boolean indicating whether the key exists.
```go
value, exists := cache.Get("key")
```

### Deleting Values
Remove a key-value pair from the cache using the `Delete()` method.
```go
cache.Delete("key")
```

### Clearing the Cache
Clear all key-value pairs from the cache using the `Clear()` method.
```go
cache.Clear()
```


## Custom Error
The package defines a custom error `ErrConversionType` for handling type conversion errors.

## Contributing
Contributions are welcome. Please submit pull requests with detailed explanations of changes.
