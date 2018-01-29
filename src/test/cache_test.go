package test

import (
	"testing"
	"strings"
	"strconv"
)

func TestMap(t *testing.T) {
	m := make(map[string]string)
	for i := 0; i < 10; i++ {
		m[strconv.Itoa(i)] = strings.Join([]string{"Map-", strconv.Itoa(i)}, "")
	}
	println(len(m))
	for k, v := range m {
		println(strings.Join([]string{k, v}, ":"))
		delete(m, k)
	}
	println(len(m))
}
