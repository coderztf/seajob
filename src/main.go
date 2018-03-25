package main

import (
	"net/http"
	"web"
	"log"
	"time"
	"spider/controller"
	"mycache"
	"task"
	"util"
	"strconv"
	"spider/entity"
	"fmt"
)

/**
@Todo 检查北上广是否重复创建缓存
 */
func main() {
	//创建任务队列
	taksCache := mycache.GetCache("task")
	//开启服务
	go func() {
		for url := range task.Task {
			_, err := mycache.Get(taksCache, url)
			if err == true {
				continue
			}
			mycache.Put(taksCache, url, 1)
			//base64解码
			url = util.Base642URL(url)
			go service(url)
		}
	}()
	go notify()
	//每隔1小时进行一次gc
	go func() {
		for {
			gc()
			time.Sleep(1 * time.Hour)
		}
	}()
	//注册查询接口
	http.HandleFunc("/list", web.ReadList)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

/**
爬虫服务
 */
func service(url string) {
	conf := "2"
	width, _ := strconv.Atoi(conf)
	if width == 0 {
		width = 1
	}
	for {
		controller.Service(url)
		time.Sleep(10000 * time.Millisecond)
	}
}

func notify() {
	for {
		log.Print("check notify")
		userCache := mycache.GetCache("user")
		locationCache := mycache.GetCache("location")
		for k, v := range *userCache {
			userInfo := v.(map[string]int)
			res := make([]entity.JobInfo, 0)
			f := false
			for lk, lv := range *locationCache {
				locationInfo := lv.(*entity.JobInfoList)
				if userInfo[lk] < len(*locationInfo) {
					f = true
					list := (*locationInfo)[userInfo[lk]:]
					log.Printf("New Information location : %s of user %s, total count:%d", lk, k, len(list))
					userInfo[lk] = len(list) + userInfo[lk]
					for _, item := range list {
						res = append(res, item)
					}
				}

			}
			if f {
				log.Println("send mail")
				//mail.SendMail(res)
			}
		}
		time.Sleep(2 * time.Hour)
	}

}

func gc() {
	locationCache := mycache.GetCache("location")
	for _, v := range *locationCache {
		list := v.(*entity.JobInfoList)
		fmt.Printf("before delete elements :%d", len(*list))
		for i, rcount, rlen := 0, 0, len(*list); i < rlen; i++ {
			j := i - rcount
			date, err := time.Parse("2006-01-02", (*list)[j].Date)
			if err == nil && date.AddDate(0, 0, 3).Before(time.Now()) {
				(*list) = append((*list)[:j], (*list)[j+1:]...)
				rcount++
			}
		}
		fmt.Printf("after delete elements :%d", len(*list))
	}
}
