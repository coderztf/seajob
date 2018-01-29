package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/entity"
)

type Parser interface {
	ConnectDocument(target string) *goquery.Selection    //连接爬取页面，并设置css选择器
	SelectorService(i int, selection *goquery.Selection) //处理经过css选择器筛选后的元素
	GetDocInfo() []entity.JobInfo                        //返回爬取信息
}
