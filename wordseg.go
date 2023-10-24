package bkit

import (
	"path"
	"regexp"
	"strings"

	"github.com/huichen/sego"
)

// 因为词典要加载到内容中,初始化时间有点长,这里用了一个全局变量,避免每次调用都初始化浪费时间
var Seg *seg

type seg struct {
	sg sego.Segmenter
}

type SegResult []string

// 需要注意, 分词使用的 github.com/huichen/sego 这个库, 使用前***必须需要初始化***
// 初始化的时候要将 data 目录中的 dict 复制到项目可以找到的地方, 并且在 NewSeg 的时候要制定 dict 目录
// 如: NewSeg(./data/dict)
func NewSeg(dictDir string) (err error) {
	if Seg != nil {
		return
	}

	Seg = &seg{}
	Seg.sg.LoadDictionary(path.Join(dictDir, "dictionary.txt"))

	return
}

// CutWords 自定义的一个快捷方法,对数据进行了一些处理,去除空白、重复、特殊字符
func (w *seg) CutWords(src string, sizes ...uint) (res SegResult) {
	if Seg == nil {
		return
	}

	var min uint = 2
	if len(sizes) > 0 {
		min = sizes[0]
	}

	// 去除准备分词的字符串中的标点和特殊字符
	src = regexp.MustCompile(`[\pP|><+=$^~]*`).ReplaceAllString(src, "")

	// 分词
	// words := w.jb.CutForSearch(src, true)
	ws := w.sg.Segment([]byte(src))
	words := sego.SegmentsToSlice(ws, true)

	res = make([]string, 0)
	m := make(map[string]bool, 0)
	// 处理只要指定长度的,并且去除重复值
	for i := 0; i < len(words); i++ {
		// 去除空白字符
		s := strings.TrimSpace(words[i])
		if len(s) == 0 {
			continue
		}

		// 处理重复
		if _, ok := m[s]; ok {
			continue
		}
		m[s] = true

		// 只要大于指定长度的词, 太短的就不要了, 比如一个字的
		if len([]rune(s)) >= int(min) {
			res = append(res, s)
		}
	}

	return
}

// String 将结果转换成字符串
func (r SegResult) String(seps ...string) string {
	// 分隔符
	sep := " "
	if len(seps) > 0 {
		sep = seps[0]
	}

	return strings.Join(r, sep)
}
