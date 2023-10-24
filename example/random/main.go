package main

import (
	"fmt"

	"github.com/broqiang/bkit"
)

func main() {
	fmt.Println("随机数", bkit.Random())

	fmt.Println("自定义随机数种子", bkit.Random(bkit.RandomOptionCustomSeed("hello23456")))

	// 基础数据使用的 []rune 类型,只要 []rune 支持的都可以,使用时注意处理特殊内容, 比如空格、换行之类的
	// 否则也会出现在随机数中
	fmt.Println("自定义中文随机数种子:", bkit.Random(bkit.RandomOptionCustomSeed(
		"你好世界我是一个中国人看看是不是可以支持中文!-+)(*&^%$#@)")))
}
