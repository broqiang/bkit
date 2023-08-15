package bkit

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 随机字符串

// 随机字符串返回类型
type RandomResult []byte

func (b RandomResult) String() string {
	return string(b)
}

// ToInt 转成 int 类型, 不保证一定成功,要保证 seed 必须是纯数字,否则此处转换失败就是 0
func (b RandomResult) Int() int {
	i, _ := strconv.Atoi(b.String())

	return i
}

// 生成随机字符串
const (
	randomLowerLetters = "abcdefghijklmnopqrstuvwxyz"
	randomUpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomNumber       = "123456789"
)

type randomStruct struct {
	seed   string
	length uint
}

type randomOption func(*randomStruct)

// RandomOptionLower 设置随机数为纯小写字母
func RandomOptionLower() randomOption {
	return func(r *randomStruct) {
		r.seed = randomLowerLetters
	}
}

// RandomOptionUpper 设置随机字符串为纯大写字母
func RandomOptionUpper() randomOption {
	return func(r *randomStruct) {
		r.seed = randomUpperLetters
	}
}

// RandomOptionLetter 设置随机字符串为大小写英文字母
func RandomOptionLetter() randomOption {
	return func(r *randomStruct) {
		r.seed = strings.Join([]string{randomUpperLetters, randomLowerLetters}, "")
	}
}

// RandomOptionNumber 设置随机字符串为大小写英文字母
func RandomOptionNumber() randomOption {
	return func(r *randomStruct) {
		r.seed = randomNumber
	}
}

// RandomOptionLength 设置随机字符串长度
func RandomOptionLength(length uint) randomOption {
	return func(r *randomStruct) {
		if length > 0 {
			r.length = length
		}
	}
}

func random(opts ...randomOption) RandomResult {
	// 设置随机数的默认种子为全部大小写英文+数字
	// 默认长度为6个字符
	rs := randomStruct{
		seed:   strings.Join([]string{randomLowerLetters, randomUpperLetters, randomNumber}, ""),
		length: 6,
	}

	for _, opt := range opts {
		opt(&rs)
	}

	var src = rand.NewSource(time.Now().UnixNano())

	var letterIdxBits uint = 6
	var letterIdxMask int64 = 1<<letterIdxBits - 1
	var letterIdxMax = 63 / letterIdxBits

	n := int(rs.length)

	seed := rs.seed

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(seed) {
			b[i] = seed[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

// Random 生成随机字符串,默认为大小写英文+数字, 长度为6
func Random(opts ...randomOption) RandomResult {
	return random(opts...)
}
