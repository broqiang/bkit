package bkit

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	// 生成默认的随机数
	fmt.Println("默认随机字符串, []byte:", Random())

	fmt.Println("默认随机字符串, string:", Random(RandomOptionLower()).String())

	fmt.Println("默认随机字符串, int:", Random().Int())

	fmt.Println("大写英文", Random(RandomOptionUpper()).String())

	fmt.Println("大小写全部英文", Random(RandomOptionLetter()).String())

	fmt.Println("10位纯数字", Random(RandomOptionNumber(), RandomOptionLength(10)).Int())

	fmt.Println("20位纯数字字符串", Random(RandomOptionNumber(), RandomOptionLength(20)).String())
}
