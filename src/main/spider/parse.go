package spider

import "github.com/PuerkitoBio/goquery"

type Parser interface {
	SelectorService(i int, selection *goquery.Selection)
	ConnectDocument(target, selector string) *goquery.Selection
	GetDocInfo() []JobInfo
}

type MySpider interface {
	DocumentParsing(url string, parser *Parser) []JobInfo
}
