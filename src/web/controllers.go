package web

import (
	"net/http"
	"mycache"
	"log"
	"spider/entity"
	"strings"
	"task"
	"util"
	"conf"
)

func ReadList(w http.ResponseWriter, r *http.Request) {
	//获取地域
	r.ParseForm()
	location := r.Form["location"][0]
	location = conf.LocationInfo[location]
	user := r.Form["user"][0]
	//加入任务队列
	task.Task <- util.URL2Base64(strings.Join([]string{"http://", "www.yingjiesheng.com/", location}, ""))
	//得到用户的已读表
	userCache := mycache.GetCache("user")
	_, exists := mycache.Get(userCache, user)
	if exists == false {
		log.Printf("%s is a new user", user)
		tmp := make(map[string]int)
		mycache.Put(userCache, user, tmp)
	}
	tmp, _ := mycache.Get(userCache, user)
	userInfo := tmp.(map[string]int)
	index := userInfo[location]
	//地域缓存
	locationCache := mycache.GetCache("location")
	temp, exists := mycache.Get(locationCache, util.URL2Location(location))
	if exists == false {
		log.Printf("%s doesn't has information \n", location)
		return
	}
	info := temp.(*entity.JobInfoList)
	list := (*info)[index:]
	//修改用户缓存
	(userInfo)[location] = len(list)+index
	log.Printf("查询新消息%d条\n", len(list))
}

