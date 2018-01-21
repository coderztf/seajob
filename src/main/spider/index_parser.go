package spider

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"main/util"
)

type IndexPaser struct {
	list []JobInfo
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
	url = util.SubString(url, strings.LastIndex(url, "/"), len(url))
	Id := util.SubString(url, strings.Index(url, "/")+1, strings.LastIndex(url, "."))
	this.list = append(this.list, JobInfo{Id, location, date, title, url})
}

func (this *IndexPaser) ConnectDocument(target, selector string) *goquery.Selection {
	doc, err := goquery.NewDocument(target)
	if err != nil {
		panic(err.Error())
		return nil
	}
	res := doc.Find(selector)
	return res
}

func (this *IndexPaser) GetDocInfo() []JobInfo{
	return this.list
}
