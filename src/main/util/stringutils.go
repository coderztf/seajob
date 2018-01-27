package util

import (
	"golang.org/x/text/transform"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
	"encoding/base64"
)

func SubString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start >= length {
		panic("start is wrong")
	}
	if end < 0 || end > length {
		panic("end is wrong")
	}
	if start == end {
		return ""
	}
	return string(rs[start:end])
}

func Gbk2Utf8(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	res, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func Title2Location(str string) (title, location string) {
	exp := regexp.MustCompile("^\\[[^\\]]+\\]")
	location = exp.FindString(str)
	title = strings.Replace(str, location, "", 1)
	location = SubString(location, 1, len(([]rune)(location))-1)
	return title, location
}

func URL2Id(str string) (url, id string) {
	url = SubString(str, strings.LastIndex(str, "/"), len(str))
	id = SubString(str, strings.Index(str, "/")+1, strings.LastIndex(str, "."))
	return url, id
}

func URL2Page(str string) (int, string) {
	exp := regexp.MustCompile("list_\\d+\\.html$")
	pageURL := exp.FindString(str)
	if pageURL == "" {
		return 0, str
	}
	str = strings.Replace(str, pageURL, "", 1)
	pageURL = strings.Replace(pageURL, "list_", "", 1)
	pageURL = strings.Replace(pageURL, ".html", "", 1)
	page, err := strconv.Atoi(pageURL)
	if err != nil {
		return 0, str
	}
	return page, str
}

func URL2Location(str string) (url string) {
	exp := regexp.MustCompile("\\w+job$")
	url = exp.FindString(str)
	url = strings.Replace(url, "job", "", 1)
	return url
}

func URL2Base64(str string) string {
	res := base64.StdEncoding.EncodeToString([]byte(str))
	return res
}

func Base642URL(str string) string {
	res, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str
	}
	return string(res)
}
