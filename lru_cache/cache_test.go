package lru_cache

import (
	"fmt"
	"testing"
)

func TestLRUCache_SetNewValues(t *testing.T) {

	lruCache, _ := NewLRUCache(3)

	for i := 0; i < 3; i++ {

		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)

		ok := lruCache.Set(Key(key), value)

		if ok {

			t.Errorf("SetNewValues error: set return true for non-existet value")
		}

		if lruCache.queue.Len() != i+1 {

			t.Errorf("SetNewValues error: invalid queue len")
		}
	}
}

func TestLRUCache_SetExistentValue(t *testing.T) {

	lruCache := makeAndFillCache(3)

	_ = lruCache.Set(Key("key_1"), "value_1")

	ok := lruCache.Set(Key("key_1"), "value_2")

	if !ok {

		t.Errorf("SetExistentValue error: Set() return not true for existent key")
	}
}

func TestLRUCache_GetNonExistentValue(t *testing.T) {

	lruCache, _ := NewLRUCache(3)

	value, ok := lruCache.Get(Key("key_1"))

	if ok {

		t.Errorf("GetNonExistentValue error: Get() return not false flag for non-existent key")
	}

	if value != nil {

		t.Errorf("GetNonExistentValue error: Get() return not nil value for non-existent key")
	}
}

func TestLRUCache_GetExistentValue(t *testing.T) {

	lruCache := makeAndFillCache(3)

	lruCache.Set(Key("key_1"), "value_1")

	value, ok := lruCache.Get(Key("key_1"))

	if !ok {

		t.Errorf("GetExistentValue error: Get() return not true flag for existent key")
	}

	if value.(string) != "value_1" {

		t.Errorf("GetNonExistentValue error: Get() return ivalid value for existent key")
	}
}

func TestLRUCache_GetQueueOrder(t *testing.T) {

	lruCache := makeAndFillCache(3)

	lruCache.Get(Key("key_1"))

	first := lruCache.queue.Front()

	if first.Value.(QueueElement).Value != "value_1" {

		t.Errorf("GetQueueOrder error: Get() does not change queue")
	}
}

func TestLRUCache_Clear(t *testing.T) {

	lruCache := makeAndFillCache(3)

	lruCache.Clear()

	if lruCache.queue.Len() != 0 {

		t.Errorf("Clear error: Clear() queue len invalid")
	}

	if len(lruCache.items) != 0 {

		t.Errorf("Clear error: Clear() items len invalid")
	}
}

func TestLRUCache_QueueOrderWithoutDuplicates(t *testing.T) {

	lruCache := makeAndFillCache(3)

	first := lruCache.queue.Front()

	if first.Value.(QueueElement).Value != "value_3" {

		t.Errorf("QueueOrderWithoutDuplicates error: invalid queue order")
	}

	last := lruCache.queue.Back()

	if last.Value.(QueueElement).Value != "value_1" {

		t.Errorf("QueueOrderWithoutDuplicates error: invalid queue order")
	}
}

func TestLRUCache_Overflow(t *testing.T) {

	lruCache := makeAndFillCache(2)

	_, ok := lruCache.Get(Key("value1"))

	if ok {

		t.Errorf("Overflow error: return old value")
	}

	if len(lruCache.items) != 2 {

		t.Errorf("Overflow error: items size greater then expected")
	}

	if lruCache.queue.Len() != 2 {

		t.Errorf("Overflow error: return old value")
	}
}

func makeAndFillCache(cap int) *LRUCache {

	lruCache, _ := NewLRUCache(cap)

	for i := 1; i <= cap; i++ {

		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)

		lruCache.Set(Key(key), value)
	}

	return lruCache
}
