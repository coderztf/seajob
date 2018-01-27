package test

import (
	"testing"
	"main/util"
	"fmt"
)

func TestConf(t *testing.T) {
	conf := util.GetConfig()
	fmt.Println(conf)
}
