package lru

import (
	"container/list"
)

type Pair struct {
	key   int
	value int
}

// Cache represents a fixed-size LRU cache for integer keys and values
type Cache struct {
	capacity int
	l        *list.List
	dict     map[int]*list.Element
}

// New returns an initialised LRU cache for the given capacity
func New(capacity int) *Cache {
	return &(Cache{capacity, list.New(), make(map[int]*list.Element)})
}

// Get returns a cached value for the given key, or -1 if the key does not exist
func (cache *Cache) Get(key int) int {
	v, ok := cache.dict[key]
	if ok == true {
		cache.l.MoveToBack(v)
		return v.Value.(*Pair).value
	}
	return -1
}

// Put inserts or updates the value for the given key.
// When the cache capacity is reached, it removes the least recently used item before inserting a new one.
func (cache *Cache) Put(key int, value int) {
	v, ok := cache.dict[key]
	if ok == true {
		v.Value.(*Pair).value = value
		cache.l.MoveToBack(v)
	} else {
		if cache.l.Len() == cache.capacity {
			delete(cache.dict, cache.l.Remove(cache.l.Front()).(*Pair).key)
		}
		cache.dict[key] = cache.l.PushBack(&Pair{key, value})
	}
}
