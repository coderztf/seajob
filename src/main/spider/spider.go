package spider

import (
	"main/mycache"
)

type JobInfo struct {
	Id       string
	Location string
	Date     string
	Title    string
	Url      string
}

var(
	spider = DefaultSpider{}
)

func getDocument(url string) []JobInfo{
	var parser *IndexPaser
	//parser必须是一个指针类型，否则会报错：method has pointer receiver
	parser = &IndexPaser{make([]JobInfo,0)}
	list := spider.DocumentParsing(url,parser)
	return list
}

func Service(url string) {
	for _, item := range getDocument(url) {
		_, ok := mycache.Get(mycache.GetFinCache(), item.Id)
		if ok {
			//缓存命中
			continue
		}
		mycache.Put(mycache.GetFinCache(), item.Id, item)
		mycache.Put(mycache.GetTodoCache(),item.Id,item)
	}
}
