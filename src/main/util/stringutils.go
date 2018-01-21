package util

import (
	"golang.org/x/text/transform"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
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
	return string(rs[start:end])
}

func Gbk2Utf8(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	res, err := ioutil.ReadAll(reader)
	if err != nil {
		return "",err
	}
	return string(res),nil
}
