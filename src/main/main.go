package main

import (
	"net/http"
	"main/web"
	"log"
	"time"
	"main/spider/controller"
	"main/mycache"
	"main/task"
	"main/util"
)

/**
@Todo 实现缓存的初始化管理
@Todo 缓存传参使用地址
@Todo 实现多用户
@Todo 根据时间控制宽度
@Todo 实现跨域调用:拦截器实现
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
			url =util.Base642URL(url)
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
	for{
		controller.Service(url)
		time.Sleep(300000 * time.Millisecond)
	}
}
