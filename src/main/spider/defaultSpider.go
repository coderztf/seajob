package spider

import "log"

type DefaultSpider struct {
}

func (defaultSpider *DefaultSpider) DocumentParsing(url string, parser *Parser) []JobInfo{
	log.Println("连接页面...")
	selector := (*parser).ConnectDocument(url,"div.box.floatl ul li")
	log.Println("解析页面信息...")
	selector.Each((*parser).SelectorService)
	return (*parser).GetDocInfo()
}