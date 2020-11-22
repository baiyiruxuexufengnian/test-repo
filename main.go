package main

import (
	"fmt"
	"github.com/baiyiruxuexufengnian/study/helper"
)

func main() {
	f := helper.NewFileHelper()
	err := f.CopyDir("test1", "test2")
	if err != nil {
		fmt.Printf("拷贝函数有问题：%v\n", err)
	} else {
		fmt.Println("成功！！")
	}
}
