package bkit

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("隐藏手机号中间4位", HiddenStr("13800138000"))

	fmt.Println("隐藏中文内容的中间4位", HiddenStr("隐藏一个中文的内容,试试看可以不可以"))

	fmt.Println("设置隐藏的开始位置", HiddenStr("13800138000", HiddenStrStart(8)))

	fmt.Println("设置隐藏的开始位置和长度", HiddenStr("13800138000", HiddenStrStart(2), HiddenStrLength(7)))

	fmt.Println("从尾部开始隐藏", HiddenStrTail("13800138000", 4))

	fmt.Println("从头部开始隐藏", HiddenStrHead("13800138000", 4))

}
