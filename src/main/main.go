package main

import (
	"main/spider"
	"net/http"
	"main/web"
	"log"
	"time"
)

func main(){
	//开启服务
	go func() {
		for{
			spider.Service("http://www.yingjiesheng.com/")
			time.Sleep(3000*time.Millisecond)
		}
	}()
	//注册查询接口
	http.HandleFunc("/list",web.ReadList)
	err :=http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal(err.Error())
	}
}
