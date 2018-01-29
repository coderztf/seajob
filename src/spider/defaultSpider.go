package spider

import (
	"log"
	"spider/entity"
	"spider/parser"
)

type DefaultSpider struct {
}

func (defaultSpider *DefaultSpider) DocumentParsing(url string, parser parser.Parser) []entity.JobInfo {
	log.Println("连接页面...")
	selector := parser.ConnectDocument(url)
	log.Println("解析页面信息...")
	selector.Each(parser.SelectorService)
	return parser.GetDocInfo()
}
