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
	if element, exist := c.items[key]; exist {
		element.Key = key
		element.Value = value

		c.queue.MoveToFront(element)
		c.items[key] = element

		return true
	}
	item := ListItem{Value: value}

	element := c.queue.PushFront(item)
	c.items[key] = element
	element.Key = key
	element.Value = value

	if c.queue.Len() > c.capacity {
		c.Clear()
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if element, exist := c.items[key]; exist {
		value := element.Value.(int)
		c.queue.MoveToFront(element)
		return value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	if element := c.queue.Back(); element != nil {
		c.queue.Remove(element)
		delete(c.items, element.Key)
	}
}
