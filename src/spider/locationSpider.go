package spider

import (
	"spider/entity"
	"log"
	"spider/parser"
)

type LocationSpider struct {
}

func (spider *LocationSpider) DocumentParsing(url string, parser parser.Parser) []entity.JobInfo {
	log.Println("连接页面...")
	selector := (parser).ConnectDocument(url)
	log.Println("解析页面信息...")
	selector.Each((parser).SelectorService)
	return (parser).GetDocInfo()
}
