package main

import (
	"sync"
)

type Cache struct {
	mu    sync.Mutex
	store map[string]string
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, found := c.store[key]
	return val, found
}
