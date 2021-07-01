package lib

import (
	"container/list"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRU) newItem(key string, value interface{}) *Item {
	return &Item{
		Key:   key,
		Value: value,
	}
}

func (c *LRU) Set(key string, value interface{}) bool {
	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		c.purge()
	}

	item := c.newItem(key, value)

	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return true
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}

func (c *LRU) Get(key string) interface{} {
	element, exists := c.items[key]

	if exists == false {
		return exists
	}

	c.queue.MoveToFront(element)

	return element.Value.(*Item).Value
}
