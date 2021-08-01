package cache

import (
	"sync"
	"time"
)

type memory struct {
	store *sync.Map
}

type memoryItem struct {
	val   interface{}
	timer *time.Timer
}

func NewMemoryCache() Cache {
	mem := memory{
		store: new(sync.Map),
	}

	return &mem
}

func (c *memory) Get(key string) (interface{}, bool) {
	item, ok := c.store.Load(key)
	if !ok {
		return nil, false
	}

	return item.(memoryItem).val, true
}

func (c *memory) Set(key string, val interface{}, timeout time.Duration) {
	timer := time.AfterFunc(timeout, func() {
		_ = c.Delete(key)

	})

	item := memoryItem{val, timer}

	c.store.Store(key, item)
}

func (c *memory) Delete(key string) error {
	c.store.Delete(key)
	return nil
}
