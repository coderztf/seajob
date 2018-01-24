package parser

import (
	"github.com/PuerkitoBio/goquery"
	"main/util"
	"fmt"
	"main/spider/entity"
)

type ProvParser struct {
	List []entity.JobInfo
}

func (this *ProvParser) SelectorService(i int, selection *goquery.Selection) {
	defer func() {
		if err := recover(); err != nil {
			html, _ := selection.Html()
			fmt.Errorf("%s \t occurred at %s", err, string(html))
		}
	}()
	if selection.HasClass("bg_0") || selection.HasClass("bg_1") {
		//有效行
		title := selection.Find("td.item1 a").Text()
		url := selection.Find("td.item1 a").AttrOr("href", "")
		date := selection.Find("td.date.cen").Text()
		//处理中文乱码
		title, _ = util.Gbk2Utf8(title)
		url, _ = util.Gbk2Utf8(url)
		date, _ = util.Gbk2Utf8(date)
		var location, id string
		title, location = util.Title2Location(title)
		url, id = util.URL2Id(url)
		this.List = append(this.List, entity.JobInfo{id, location, date, title, url})
	}
}

func (this *ProvParser) ConnectDocument(target string) *goquery.Selection {
	doc, err := goquery.NewDocument(target)
	if err != nil {
		panic(err.Error())
		return nil
	}
	res := doc.Find("table tbody tr")
	return res
}

func (this *ProvParser) GetDocInfo() []entity.JobInfo {
	return this.List
}
