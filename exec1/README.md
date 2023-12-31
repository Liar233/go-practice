# LRU-cache

## Цель:

Реализовать [LRU-cache](https://blog.skillfactory.ru/glossary/lru/) на основе [двухсвязного списка](https://pkg.go.dev/container/list@go1.21.1)

## Задание:

Необходимо реализовать LRU-кэш на основе двусвязного списка.

1. Реализовать interface Cache

```go
type Key string

type Cache interface {
    Set(key Key, value interface{}) bool
    Get(key Key) (interface{}, bool)
    Clear()
}
```

2. Структура Cache

```go
type LRUCache struct {
	// количество сохраняемых в кэше элементов
	cap int
	// хеш-таблица отображающая ключ (строка) на элемент очереди 
	items    map[Key]*list.Element
	// последних используемых элементов 
	queue    *list.List
}
```

3. 1. Алгоритм работы кэша:

Set:

- если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди
- если элемента нет в словаре, то добавить в словарь и в начало очереди (при этом, если размер очереди больше ёмкости кэша, то необходимо удалить последний элемент из очереди и его значение из словаря)
- возвращаемое значение - флаг, присутствовал ли элемент в кэше

Get:

- eсли элемент присутствует в словаре, то переместить элемент в начало очереди и вернуть его значение и true
- если элемента нет в словаре, то вернуть nil и false
