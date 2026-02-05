package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, exists := c.items[key]; exists {
		item.Value = cacheEntry{key: key, value: value}
		c.queue.MoveToFront(item)
		return true
	}

	entry := c.queue.PushFront(cacheEntry{key: key, value: value})
	c.items[key] = entry

	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		if lastItem != nil {
			for k, item := range c.items {
				if item == lastItem {
					delete(c.items, k)
					break
				}
			}
			c.queue.Remove(lastItem)
		}
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, exists := c.items[key]; exists {
		c.queue.MoveToFront(item)
		if entry, ok := item.Value.(cacheEntry); ok {
			return entry.value, true
		}
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

type cacheEntry struct {
	key   Key
	value interface{}
}
