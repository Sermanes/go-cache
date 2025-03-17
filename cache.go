package cache

import (
	"errors"
	"sync"
)

var (
	// The line `ErrConversionType = errors.New("error converting type")` is defining a new error variable
	// named `ErrConversionType` with a specific error message "error converting type". This error
	// variable can be used to represent a specific type of error related to type conversion within the
	// cache package. It allows the code to return this specific error when there is an issue with
	// converting types during operations in the cache.
	ErrConversionType = errors.New("error converting type")
)

// The Cache type in Go represents a data structure with a map for storing key-value pairs and a mutex
// for synchronization.
// @property store - The `store` property in the `Cache` struct is a map that stores key-value pairs
// where the keys are strings and the values can be of any type.
// @property mu - The `mu` property in the `Cache` struct is a `sync.RWMutex` type. It is used to
// provide locking mechanisms to control access to the `store` map in a concurrent-safe manner. The
// `RWMutex` allows multiple readers or a single writer to access the map
type Cache struct {
	store map[string]any
	mu    sync.RWMutex
}

var (
	instance *Cache
	once     sync.Once
)

func GetInstance() *Cache {
	once.Do(func() {
		instance = &Cache{
			store: make(map[string]any),
		}
	})
	return instance
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.store[key]
	return value, exists
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store = make(map[string]any)
}
