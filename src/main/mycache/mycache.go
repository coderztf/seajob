package mycache

import (
	"sync"
	"main/util"
)

type MyCache map[string]interface{}

var once sync.Once

var sys MyCache

var mux sync.Mutex

func init() {
	once.Do(func() {
		sys = MyCache{}
	})
}

func GetCache(name string) *MyCache{
	cache,err := Get(&sys,name)
	if err == false{
		cache =&MyCache{}
		Put(&sys,name,cache)
	}
	return cache.(*MyCache)
}

func Put(cache *MyCache, key string, value interface{}) {
	mux.Lock()
	defer mux.Unlock()
	key = util.URL2Base64(key)
	(*cache)[key] = value
}

func Get(cache *MyCache, key string) (interface{}, bool) {
	mux.Lock()
	defer mux.Unlock()
	key = util.URL2Base64(key)
	elem, ok := (*cache)[key]
	return elem, ok
}

func Remove(cache *MyCache, key string) {
	mux.Lock()
	key = util.URL2Base64(key)
	delete(*cache, key)
	mux.Unlock()
}