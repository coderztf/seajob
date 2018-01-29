package parser

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"spider/entity"
	"util"
	"strings"
	"sync"
)

type LocationParser struct {
	locationInfo map[string]string
	mux          sync.Mutex
}

func (parser *LocationParser) SelectorService(i int, selection *goquery.Selection) {
	//锁定信息
	parser.mux.Lock()
	defer func() {
		if err := recover(); err != nil {
			html, _ := selection.Html()
			fmt.Errorf("%ss \t occurred at %s", err, string(html))
		}
		parser.mux.Unlock()
	}()
	if selection.HasClass("sideMenu") {
		//sideMenu
		selection.Find("li a").Each(func(i int, selection *goquery.Selection) {
			location := selection.Text()
			location, _ = util.Gbk2Utf8(location)
			url := selection.AttrOr("href", "")
			parser.locationInfo[location] = strings.Replace(url, "/", "", -1)
		})
	}
	if selection.HasClass("pubMenu") {
		//pubMenu
		selection.Find("li a:first").Each(func(i int, selection *goquery.Selection) {
			location := selection.Text()
			location, _ = util.Gbk2Utf8(location)
			url := selection.AttrOr("href", "")
			parser.locationInfo[location] = strings.Replace(url, "/", "", -1)
		})
	}
}

func (parser *LocationParser) ConnectDocument(target string) *goquery.Selection {
	doc, err := goquery.NewDocument(target)
	if err != nil {
		panic(err.Error())
		return nil
	}
	res := doc.Find("div.menu ul")
	return res
}

/**
区域信息爬取不需要保存职位信息
 */
func (parser *LocationParser) GetDocInfo() []entity.JobInfo {
	return nil
}

func (parser *LocationParser) GetLocationInfo() map[string]string {
	parser.mux.Lock()
	defer parser.mux.Unlock()
	return parser.locationInfo
}

func GetLocationParser() *LocationParser {
	parser := LocationParser{locationInfo: make(map[string]string)}
	return &parser
}
