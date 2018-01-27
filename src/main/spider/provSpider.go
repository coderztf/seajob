package spider

import (
	"main/spider/entity"
	"main/util"
	"strings"
	"strconv"
	"sync"
	"fmt"
)

type ProvSpider struct {
	page int
}

func (provSpider *ProvSpider) DocumentParsing(url string, parser Parser) []entity.JobInfo {
	list := make([]entity.JobInfo, 0)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	if page, nonePage := util.URL2Page(url); page == 0 && page < 5 {
		//开启三个线程同时爬取
		for i := 1; i <= 4; i++ {
			pageURL := strings.Join([]string{nonePage, strings.Join([]string{"/list_", strconv.Itoa(page + i)}, ""), ".html"}, "")
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
		selector := parser.ConnectDocument(url)
		selector.Each(parser.SelectorService)
		list = parser.GetDocInfo()
	}
	wg.Wait()
	fmt.Printf("url : %s \t has %d infos\n", url, len(list))
	return list
}
