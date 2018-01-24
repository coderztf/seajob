package mycache

import (
	"main/spider/entity"
	"sync"
)

/**
	缓存数据结构
 */
type CacheInfo struct {
	Name     string
	Todo     map[string]entity.JobInfo
	Fin map[string]entity.JobInfo
	Lock sync.Mutex
}

func InitCacheInfo(name string) CacheInfo {
	cache := CacheInfo{Name: name}
	cache.Todo = make(map[string]entity.JobInfo)
	cache.Fin = make(map[string]entity.JobInfo)
	return cache
}