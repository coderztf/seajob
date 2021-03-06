package controller

import (
	"mycache"
	"spider"
	"util"
	"log"
	"spider/entity"
	"spider/parser"
)

var (
	defaultSpider = spider.DefaultSpider{}
	provSpider    = spider.ProvSpider{}
)

func getDocument(url string) []entity.JobInfo {
	//parser必须是一个指针类型，否则会报错：method has pointer receiver
	//parser := &parser.IndexPaser{make([]entity.JobInfo,0)}
	//list := defaultSpider.DocumentParsing(url,parser)
	parser := &(parser.ProvParser{List: make([]entity.JobInfo, 0)})
	list := provSpider.DocumentParsing(url, parser)
	return list
}

func Service(url string) {
	location := util.URL2Location(url)
	locationCache := mycache.GetCache("location")
	cache, exists := mycache.Get(locationCache, location)
	if exists == false {
		mycache.Put(locationCache, location, entity.InitJobInfoList())
		log.Printf("%s created location cache\n", location)
		cache, _ = mycache.Get(locationCache, location)
	}
	info := cache.(*entity.JobInfoList)
	for _, item := range getDocument(url) {
		_, ok := info.Index(item)
		if ok {
			//缓存命中
			continue
		}
		info.Add(item)
	}
}
