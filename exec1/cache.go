package main

import "container/list"

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

func NewLRUCache(cap int) *LRUCache {

	return &LRUCache{
		cap:   cap,
		items: make(map[Key]*list.Element, cap),
		queue: list.New(),
	}
}

func (LRU *LRUCache) Set(key Key, value interface{}) bool {

	el, ok := LRU.items[key]
	newEl := LRU.queue.PushFront(value)
	LRU.items[key] = newEl

	if ok {

		LRU.queue.Remove(el)

		return true
	}

	if LRU.queue.Len() > LRU.cap {

		last := LRU.queue.Back()

		for k, val := range LRU.items {

			if val == last {

				delete(LRU.items, k)
				break
			}
		}

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

	return val.Value, true
}

func (LRU *LRUCache) Clear() {

	LRU.items = make(map[Key]*list.Element, LRU.cap)
	LRU.queue = list.New()
}
