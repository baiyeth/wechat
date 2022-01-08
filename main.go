package main

import (
	"fmt"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func main() {
	data, err := wechat.Code2Session("testCode2Session", "", "")
	fmt.Printf("call Code2Session:%v error:%v", data, err)
}
