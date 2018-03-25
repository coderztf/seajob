package util

import (
	"os"
	"log"
	"bufio"
	"io"
	"strings"
	"fmt"
	"path/filepath"
)

type configInfo struct {
	name    string
	confMap map[string]string
}

type config []configInfo

var (
	conf         config
	LocationInfo map[string]string
)

/**
获得配置信息实体
 */
func getConfig() *config {
	return &conf
}

func init2() {
	conf = make([]configInfo, 0)
	current, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	parrent := SubString(current, 0, strings.LastIndex(current, "/"))
	file, err := os.Open(strings.Join([]string{parrent, "/conf/config.conf"}, ""))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		bline, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err.Error())
			return
		}
		//获得一行数据
		str := strings.TrimSpace(string(bline))
		if deleteComment(&str); len(strings.TrimSpace(str)) == 0 {
			//注释行或者空行
			continue
		}

		name := confName(str)
		if name != "" {
			//单元行
			info := configInfo{name: name, confMap: make(map[string]string)}
			conf = append(conf, info) //添加元素
			continue
		}
		//配置信息行
		key, value := info(str)
		if len(conf) == 0 {
			//创建初始信息单元
			conf = append(conf, configInfo{name: "", confMap: (make(map[string]string))})
		}
		current := &(conf[len(conf)-1])
		current.confMap[key] = value
	}
}

func confName(str string) (name string) {
	n1 := strings.Index(str, "[")
	n2 := strings.Index(str, "]")
	if n1 > -1 && n2 > -1 && n2 > n1 {
		name = strings.TrimSpace(str[n1+1:n2])
	}
	return name
}

func info(str string) (key, value string) {
	strs := strings.Split(str, "=")
	if len(strs) == 0 {
		return key, value
	}
	key = strs[0]
	value = strings.Join(strs[1:], "")
	return strings.TrimSpace(key), strings.TrimSpace(value)
}

func deleteComment(str *string) {
	index := strings.Index(*str, "#")
	if index >= 0 {
		*str = SubString(*str, 0, index)
	}
	index = strings.Index(*str, "//")
	if index >= 0 {
		*str = SubString(*str, 0, index)
	}
}

func (conf config) String() string {
	var res string
	for _, item := range conf {
		res = strings.Join([]string{res, fmt.Sprintf("[%s]\n", item.name)}, "")
		for key, value := range item.confMap {
			res = strings.Join([]string{res, fmt.Sprintf("\t%s=%s\n", key, value)}, "")
		}
	}
	return res
}

func (conf *config) Get(name, key string) string {
	for _, item := range *conf {
		if item.name == name {
			return item.confMap[key]
		}
	}
	return ""
}