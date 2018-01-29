package test

import (
	"testing"
	"util"
	"fmt"
)

func TestConf(t *testing.T) {
	conf := util.GetConfig()
	fmt.Println(conf)
}
