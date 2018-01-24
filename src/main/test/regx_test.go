package test

import (
	"testing"
	"main/util"
	"fmt"
)

func TestRegx(t *testing.T){
	title,location :=util.Title2Location("[浙江]杭州一隅千象科技有限公司图像算法工程师")
	fmt.Println(title)
	fmt.Println(location)
}
