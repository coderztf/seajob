package test

import (
	"testing"
	"spider"
	"spider/parser"
)

func TestSpider(t *testing.T) {
	parser := &(parser.LocationParser{})
	spider := spider.LocationSpider{}
	spider.DocumentParsing("http://www.yingjiesheng.com", parser)
}
