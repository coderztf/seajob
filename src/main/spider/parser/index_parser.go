package parser

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"main/util"
	"main/spider/entity"
)

type IndexPaser struct {
	List []entity.JobInfo
}

func (this *IndexPaser) SelectorService(i int, selection *goquery.Selection) {
	if selection.HasClass("warn") || selection.HasClass("space") || selection.HasClass("more") {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			html, _ := selection.Html()
			fmt.Errorf("%s \t occurred at %s", err, string(html))
		}
	}()
	url, _ := selection.Find("a").Attr("href")
	date := selection.Find("span").Last().Text()
	location := selection.Find("a span").Text()
	selection.Find("a span").Remove()
	title := selection.Find("a").Text()
	//处理字符串
	date, _ = util.Gbk2Utf8(date)
	date = util.SubString(date, 1, len(date)-1)
	title, _ = util.Gbk2Utf8(title)
	location, _ = util.Gbk2Utf8(location)
	url, _ = util.Gbk2Utf8(url)
	var Id string
	url, Id = util.URL2Id(url)
	this.List = append(this.List, entity.JobInfo{Id, location, date, title, url})
}

func (this *IndexPaser) ConnectDocument(target string) *goquery.Selection {
	doc, err := goquery.NewDocument(target)
	if err != nil {
		panic(err.Error())
		return nil
	}
	res := doc.Find("div.box.floatl ul li")
	return res
}

func (this *IndexPaser) GetDocInfo() []entity.JobInfo {
	return this.List
}
