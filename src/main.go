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
)

/**
@Todo 实现跨域调用:拦截器实现
@Todo 实现地域的智能匹配
@Todo 优化区域爬虫逻辑，根据日期，自动开启多线程爬取
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
	conf := util.GetConfig().Get("service", "width")
	width, _ := strconv.Atoi(conf)
	if width == 0 {
		width = 1
	}
	for i := 0; i < width; i++ {
		controller.Service(url)
		time.Sleep(10000 * time.Millisecond)
	}
}
