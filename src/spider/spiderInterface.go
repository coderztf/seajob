package spider

import (
	"spider/entity"
	"spider/parser"
)

type MySpider interface {
	DocumentParsing(url string, parser parser.Parser) []entity.JobInfo //执行爬虫
}
