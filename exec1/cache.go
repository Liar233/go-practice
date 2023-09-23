package main

import (
	"container/list"
	"errors"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type LRUCache struct {
	// количество сохраняемых в кэше элементов
	cap int
	// хеш-таблица отображающая ключ (строка) на элемент очереди
	items map[Key]*list.Element
	// последних используемых элементов
	queue *list.List
}

type QueueElement struct {
	Key   Key
	Value interface{}
}

func NewLRUCache(cap int) (*LRUCache, error) {

	if cap <= 0 {

		return nil, errors.New("capacity less less then one")
	}

	return &LRUCache{
		cap:   cap,
		items: make(map[Key]*list.Element, cap),
		queue: list.New(),
	}, nil
}

func (LRU *LRUCache) Set(key Key, value interface{}) bool {

	queueElement := QueueElement{
		Key:   key,
		Value: value,
	}

	newElement := LRU.queue.PushFront(queueElement)

	element, ok := LRU.items[key]
	LRU.items[key] = newElement

	if ok {

		LRU.queue.Remove(element)

		return true
	}

	if LRU.queue.Len() > LRU.cap {

		last := LRU.queue.Back()

		delete(LRU.items, last.Value.(QueueElement).Key)

		LRU.queue.Remove(last)
	}

	return false
}

func (LRU *LRUCache) Get(key Key) (interface{}, bool) {

	val, ok := LRU.items[key]

	if ok != true {

		return nil, false
	}

	if val != nil {

		LRU.queue.MoveToFront(val)
	}

	return val.Value.(QueueElement).Value, true
}

func (LRU *LRUCache) Clear() {

	LRU.items = make(map[Key]*list.Element, LRU.cap)
	LRU.queue = list.New()
}
