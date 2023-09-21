package main

import (
	"fmt"
	"testing"
)

func TestLRUCache_SetNewValues(t *testing.T) {

	var ok bool

	lurCache := NewLRUCache(3)

	for i := 0; i < 3; i++ {

		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)

		ok = lurCache.Set(Key(key), value)

		if ok != false {

			t.Fatalf("SetNewValues error: set return true for non-existet value")
		}

		if lurCache.queue.Len() != i+1 {

			t.Fatalf("SetNewValues error: invalid queue len")
		}
	}
}

func TestLRUCache_SetExistentValue(t *testing.T) {

	var ok bool

	lurCache := NewLRUCache(3)

	_ = lurCache.Set(Key("key_1"), "value_1")

	ok = lurCache.Set(Key("key_1"), "value_2")

	if ok != true {

		t.Fatalf("SetExistentValue error: Set() return not true for existent key")
	}
}

func TestLRUCache_GetNonExistentValue(t *testing.T) {

	lurCache := NewLRUCache(3)

	value, ok := lurCache.Get(Key("key_1"))

	if ok != false {

		t.Fatalf("GetNonExistentValue error: Get() return not false flag for non-existent key")
	}

	if value != nil {

		t.Fatalf("GetNonExistentValue error: Get() return not nil value for non-existent key")
	}
}

func TestLRUCache_GetExistentValue(t *testing.T) {

	lurCache := NewLRUCache(3)

	lurCache.Set(Key("key_1"), "value_1")

	value, ok := lurCache.Get(Key("key_1"))

	if ok != true {

		t.Fatalf("GetExistentValue error: Get() return not true flag for existent key")
	}

	if value.(string) != "value_1" {

		t.Fatalf("GetNonExistentValue error: Get() return ivalid value for existent key")
	}
}

func TestLRUCache_GetQueueOrder(t *testing.T) {

	lurCache := NewLRUCache(3)

	lurCache.Set(Key("key_1"), "value_1")
	lurCache.Set(Key("key_2"), "value_2")
	lurCache.Set(Key("key_3"), "value_3")

	lurCache.Get(Key("key_1"))

	first := lurCache.queue.Front()

	if first.Value.(string) != "value_1" {

		t.Fatalf("GetQueueOrder error: Get() does not change queue")
	}
}

func TestLRUCache_Clear(t *testing.T) {

	lurCache := NewLRUCache(3)

	lurCache.Set(Key("key_1"), "value_1")
	lurCache.Set(Key("key_2"), "value_2")
	lurCache.Set(Key("key_3"), "value_3")

	lurCache.Clear()

	if lurCache.queue.Len() != 0 {

		t.Fatalf("Clear error: Clear() queue len invalid")
	}

	if len(lurCache.items) != 0 {

		t.Fatalf("Clear error: Clear() items len invalid")
	}
}

func TestLRUCache_QueueOrderWithoutDuplicates(t *testing.T) {

	lurCache := NewLRUCache(3)

	lurCache.Set(Key("key_1"), "value_1")
	lurCache.Set(Key("key_2"), "value_2")
	lurCache.Set(Key("key_3"), "value_3")

	first := lurCache.queue.Front()

	if first.Value.(string) != "value_3" {

		t.Fatalf("QueueOrderWithoutDuplicates error: invalid queue order")
	}

	last := lurCache.queue.Back()

	if last.Value.(string) != "value_1" {

		t.Fatalf("QueueOrderWithoutDuplicates error: invalid queue order")
	}
}

func TestLURCache_Overflow(t *testing.T) {

	lurCache := NewLRUCache(2)

	lurCache.Set(Key("key_1"), "value_1")
	lurCache.Set(Key("key_2"), "value_2")
	lurCache.Set(Key("key_3"), "value_3")

	_, ok := lurCache.Get(Key("value1"))

	if ok != false {

		t.Fatalf("Overflow error: return old value")
	}

	if len(lurCache.items) != 2 {

		t.Fatalf("Overflow error: items size greater then expected")
	}

	if lurCache.queue.Len() != 2 {

		t.Fatalf("Overflow error: return old value")
	}
}
