package conf

import (
	parser2 "spider/parser"
	spider2 "spider"
)

var LocationInfo map[string]string

func init() {
	go func() {
		parser := parser2.GetLocationParser()
		spider := spider2.LocationSpider{}
		spider.DocumentParsing("http://www.yingjiesheng.com", parser)
		LocationInfo = parser.GetLocationInfo()
	}()
}
