package bkit

import (
	"fmt"
	"testing"
)

func TestWordSeg(t *testing.T) {
	// 必须初始化
	err := NewSeg("./data/dict")
	if err != nil {
		t.Error(err)
		return
	}

	result := Seg.CutWords("初始化的时候要将data目录中的dict复制到项目可以找到的地方复制到项目可以找到的地方复制到项目可以找到的地方")
	fmt.Println("分词结果: ", result)
	fmt.Println("分词结果是字符串: ", result.String())

	result = Seg.CutWords("小明硕士毕业于清华大学，后在中国科学院计算所深造", 4)
	fmt.Println("分词指定词长度: ", result)

}
