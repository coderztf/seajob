package controller

import (
	"main/mycache"
	"main/spider/parser"
	"main/spider/entity"
	"main/spider"
	"main/util"
	"log"
)

var (
	defaultSpider = spider.DefaultSpider{}
	provSpider    = spider.ProvSpider{}
)

func getDocument(url string) []entity.JobInfo {
	//parser必须是一个指针类型，否则会报错：method has pointer receiver
	//parser := &parser.IndexPaser{make([]entity.JobInfo,0)}
	//list := defaultSpider.DocumentParsing(url,parser)
	parser := &parser.ProvParser{make([]entity.JobInfo, 0)}
	list := provSpider.DocumentParsing(url, parser)
	return list
}

/**
@TODO 缓存加锁
 */
func Service(url string) {
	location := util.URL2Location(url)
	locationCache := mycache.GetCache("location")
	cache, exists := mycache.Get(locationCache, location)
	if exists == false {
		mycache.Put(locationCache, location, mycache.InitCacheInfo(location))
		log.Printf("%s created location cache\n",location)
		cache, _ = mycache.Get(locationCache, location)
	}
	info := cache.(mycache.CacheInfo)
	info.Lock.Lock()
	for _, item := range getDocument(url) {
		_, ok := info.Fin[item.Id]
		if ok {
			//缓存命中
			continue
		}
		info.Fin[item.Id] = item
		info.Todo[item.Id] = item
	}
	info.Lock.Unlock()
	//写回缓存
	mycache.Put(locationCache,location,info)
}
