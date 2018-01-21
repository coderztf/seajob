package mycache

import (
	"sync"
)

type MyCache map[string]interface{}

var once sync.Once

var cache MyCache

var todo MyCache

var mux sync.Mutex

func init() {
	once.Do(func() {
		cache = MyCache{}
		todo = MyCache{}
	})
}

func GetTodoCache() MyCache {
	return todo
}

func GetFinCache() MyCache {
	return cache
}

func Put(cache MyCache, key string, value interface{}) {
	mux.Lock()
	defer mux.Unlock()
	cache[key] = value
}

func Get(cache MyCache, key string) (interface{}, bool) {
	mux.Lock()
	defer mux.Unlock()
	elem, ok := cache[key]
	return elem, ok
}

func Remove(cache MyCache, key string) {
	mux.Lock()
	delete(cache, key)
	mux.Unlock()
}

func EachItem(cache MyCache, f func(key string, value interface{})) {
	for key, value := range cache {
		f(key, value)
	}
}
