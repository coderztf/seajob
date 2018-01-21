package web

import (
	"net/http"
	"main/mycache"
	"main/spider"
	"encoding/json"
	"log"
)


func ReadList(w http.ResponseWriter, r *http.Request){
	list := make([]spider.JobInfo,0)
	mycache.EachItem(mycache.GetTodoCache(), func(key string, value interface{}) {
		list = append(list,value.(spider.JobInfo))
		mycache.Remove(mycache.GetTodoCache(),key)
	})
	log.Printf("查询新消息%d条\n",len(list))
	//输出招聘信息
	json,err := json.Marshal(list)
	if err != nil{
		http.Error(w,err.Error(),500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}