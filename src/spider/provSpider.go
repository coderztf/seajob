package spider

import (
	"util"
	"strings"
	"strconv"
	"sync"
	"fmt"
	"spider/entity"
	"spider/parser"
)

type ProvSpider struct {
	page int
}

func (provSpider *ProvSpider) DocumentParsing(url string, parser parser.Parser) []entity.JobInfo {
	list := make([]entity.JobInfo, 0)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	if page, nonePage := util.URL2Page(url); page == 0 && page < 5 {
		//多线程并发爬取信息
		widthStr := "4"
		width, err := strconv.Atoi(widthStr)
		if err != nil {
			//默认单线程
			width = 1
		}
		for i := 1; i <= width; i++ {
			var pageURL string
			if strings.Contains(url, "beijing") || strings.Contains(url, "shanghai") || strings.Contains(url, "guangzhou") {
				pageURL = strings.Join([]string{nonePage, strings.Join([]string{"-morejob-", strconv.Itoa(page + i)}, ""), ".html"}, "")
			} else {
				pageURL = strings.Join([]string{nonePage, strings.Join([]string{"/list_", strconv.Itoa(page + i)}, ""), ".html"}, "")
			}
			wg.Add(1)
			go func() {
				info := provSpider.DocumentParsing(pageURL, parser)
				lock.Lock()
				list = append(list, info...)
				lock.Unlock()
				wg.Done()
			}()
		}
	} else {
		selector := (parser).ConnectDocument(url)
		selector.Each((parser).SelectorService)
		list = (parser).GetDocInfo()
	}
	wg.Wait()
	fmt.Printf("url : %s \t has %d infos\n", url, len(list))
	return list
}
