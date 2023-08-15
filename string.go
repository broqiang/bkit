package bkit

import (
	"strings"
)

// 字符串替换
type hiddenStrConfig struct {
	start  int
	length int
	symbol string
}

type hiddenStrOption func(*hiddenStrConfig)

// HiddenStrStart 设置替换的开始位置,默认是 3
func HiddenStrStart(start uint) hiddenStrOption {
	return func(rss *hiddenStrConfig) {
		rss.start = int(start)
	}
}

// HiddenStrLength 设置替换的长度, 默认是 4
func HiddenStrLength(length uint) hiddenStrOption {
	return func(rss *hiddenStrConfig) {
		rss.length = int(length)
	}
}

// HiddenStrSymbol 设置隐藏时的替换内容, 默认是 *
func HiddenStrSymbol(symbol string) hiddenStrOption {
	return func(rss *hiddenStrConfig) {
		rss.symbol = symbol
	}
}

// HiddenStr 字符串替换,用于隐藏部分内容
// 默认配置是隐藏手机中间4位, 如: 138****8000
func HiddenStr(origin string, opts ...hiddenStrOption) string {
	s := []rune(origin)

	rss := hiddenStrConfig{
		start:  3,
		length: 4,
		symbol: "*",
	}

	for _, opt := range opts {
		opt(&rss)
	}

	// 如果字符串的长度小于开始位置+长度,就什么都不处理,直接返回原始字符串
	length := len(s)

	// 如果传入字符串的长度小于替换的开始位置,直接返回原始字符串(继续执行会出错,越界)
	if length < rss.start {
		return string(s)
	}

	// 如果字符串长度小于开始位置+替换长度, 就将替换长度替换为可以替换的剩余长度
	if length < (rss.start + rss.length) {
		rss.length = length - rss.start
	}

	symbolRune := []rune(strings.Repeat(rss.symbol, rss.length))

	copy(s[rss.start:rss.start+rss.length], symbolRune)

	return string(s)
}

// HiddenStrTail 尾部开始隐藏
// Example: HiddenStrTail("13800138000", 4)
// Out: 1380013****
func HiddenStrTail(origin string, length int, opts ...hiddenStrOption) string {
	l := len(origin)

	start := l - length

	// 如果字符串的长度小于需要替换的长度, 就将全部字符串都替换
	if l < length {
		start = 0
		length = l
	}

	// 设置开始位置和替换长度
	opts = append(opts, HiddenStrStart(uint(start)), HiddenStrLength(uint(length)))

	return HiddenStr(origin, opts...)
}

// HiddenStrHead 从头部开始替换
// Example: ReplaceStrTail("13800138000", 4)
// Out: ****0138000
func HiddenStrHead(origin string, length int, opts ...hiddenStrOption) string {
	l := len(origin)

	// 如果字符串的长度小于需要替换的长度, 就将全部字符串都替换
	if l < length {
		// 修改 length 未字符串的长度,防止越界
		length = l
	}

	// 设置开始位置和替换长度
	opts = append(opts, HiddenStrStart(uint(0)), HiddenStrLength(uint(length)))

	return HiddenStr(origin, opts...)
}
