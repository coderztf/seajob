package spider

import (
	"github.com/PuerkitoBio/goquery"
	"main/spider/entity"
)

type Parser interface {
	SelectorService(i int, selection *goquery.Selection)
	ConnectDocument(target string) *goquery.Selection
	GetDocInfo() []entity.JobInfo
}

type MySpider interface {
	DocumentParsing(url string, parser *Parser) []entity.JobInfo
}
