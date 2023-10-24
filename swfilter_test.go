package bkit

import (
	"fmt"
	"testing"
)

func TestSWFilter(t *testing.T) {
	filter := SWFilter()

	afterFiltration, keywords, exist := filter.Filter("随便一些内容用于测试炸药我有炸药这是正常内容摸胸")
	fmt.Printf("替换后的内容: %s\n找到的关键词: %v\n是否存在敏感词:%t\n\n", afterFiltration, keywords, exist)
	/*
		存在敏感词,输出:
			替换后的内容: 随便一些内容用于测试**我有**这是正常内容**
			找到的关键词: [炸药 摸胸]
			是否存在敏感词:true
	*/

	afterFiltration, keywords, exist = filter.Filter("这段话不包含敏感词")
	fmt.Printf("替换后的内容: %s\n找到的关键词: %v\n是否存在敏感词:%t\n\n", afterFiltration, keywords, exist)
	/*
		不存在敏感词, 输出:
			替换后的内容: 这段话不包含敏感词
			找到的关键词: []
			是否存在敏感词:false
	*/

	// 链式操作
	afterFiltration, keywords, exist = SWFilter().Filter("测试下链式操作")
	fmt.Printf("替换后的内容: %s\n找到的关键词: %v\n是否存在敏感词:%t\n\n", afterFiltration, keywords, exist)
}
