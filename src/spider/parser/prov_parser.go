package parser

import (
	"github.com/PuerkitoBio/goquery"
	"util"
	"fmt"
	"spider/entity"
	"time"
)

type ProvParser struct {
	List []entity.JobInfo
}

func (this *ProvParser) SelectorService(i int, selection *goquery.Selection) {
	defer func() {
		if err := recover(); err != nil {
			html, _ := selection.Html()
			fmt.Errorf("%ss \t occurred at %s", err, string(html))
		}
	}()
	if selection.HasClass("bg_0") || selection.HasClass("bg_1") || selection.HasClass("tr_list") {
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
		_, id = util.URL2Id(url)
		t1, err := time.Parse("2006-01-02", date)
		if err == nil && t1.AddDate(0, 0, 3).After(time.Now()) {
			this.List = append(this.List, entity.JobInfo{id, location, date, title, url})
		}
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
