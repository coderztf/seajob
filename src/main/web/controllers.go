package web

import (
	"net/http"
	"main/mycache"
	"encoding/json"
	"log"
	"main/spider/entity"
	"strings"
	"main/task"
	"main/util"
)


func ReadList(w http.ResponseWriter, r *http.Request){
	//获取地域
	r.ParseForm()
	location := r.Form["location"][0]
	//加入任务队列
	task.Task <- util.URL2Base64(strings.Join([]string{"http://","www.yingjiesheng.com/",location,"job"},""))
	//地域缓存->todo缓存
	locationCache := mycache.GetCache("location")
	temp,exists :=mycache.Get(locationCache,location)
	if exists == false{
		log.Printf("%s doesn't has information \n",location)
		return
	}
	todo := temp.(mycache.CacheInfo).Todo
	list := make([]entity.JobInfo,0)
	for key,value := range todo{
		list = append(list,value)
		delete(todo,key)
	}
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