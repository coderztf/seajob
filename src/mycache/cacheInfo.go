package mycache

import (
	"sync"
	"spider/entity"
)

/**
	缓存数据结构
 */
type CacheJobInfo2 struct {
	Todo map[string]entity.JobInfo //新增记录
	Fin  map[string]entity.JobInfo //已处理记录
	Lock sync.Mutex
}

/**
	获得一个初始的缓存块
 */
func InitCacheJobInfo() CacheJobInfo2 {
	cache := CacheJobInfo2{Todo: make(map[string]entity.JobInfo), Fin: make(map[string]entity.JobInfo)}
	return cache
}
