package bkit

// 用于关键词替换
// 复制自: https://github.com/zeromicro/go-zero/blob/master/core/stringx/trie.go#L77

import (
	"fmt"
	"reflect"
	"strconv"
)

const defaultMask = '*'

var defaultTrie Trie

type (
	// TrieOption defines the method to customize a Trie.
	TrieOption func(trie *trieNode)

	// A Trie is a tree implementation that used to find elements rapidly.
	Trie interface {
		Filter(text string) (string, []string, bool)
		FindKeywords(text string) []string
	}

	trieNode struct {
		node
		mask rune
	}

	scope struct {
		start int
		stop  int
	}
)

// 使用方式:
// filter := bkit.SWFilter()
// string, []string, bool = filter.Filter(m.Title)

func SWFilter() Trie {
	if defaultTrie != nil {
		return defaultTrie
	}

	pLen := len(politics)
	iLen := len(illegal)
	prLen := len(pron)

	words := make([]string, pLen+iLen+prLen)
	copy(words[0:pLen], politics)
	copy(words[pLen:pLen+iLen], illegal)
	copy(words[pLen+iLen:], pron)

	t := NewTrie(words)

	defaultTrie = t

	return defaultTrie
}

// NewTrie returns a Trie.
func NewTrie(words []string, opts ...TrieOption) Trie {
	n := new(trieNode)

	for _, opt := range opts {
		opt(n)
	}
	if n.mask == 0 {
		n.mask = defaultMask
	}
	for _, word := range words {
		n.add(word)
	}

	n.build()

	return n
}

func (n *trieNode) Filter(text string) (sentence string, keywords []string, found bool) {
	chars := []rune(text)
	if len(chars) == 0 {
		return text, nil, false
	}

	scopes := n.find(chars)
	keywords = n.collectKeywords(chars, scopes)

	for _, match := range scopes {
		// we don't care about overlaps, not bringing a performance improvement
		n.replaceWithAsterisk(chars, match.start, match.stop)
	}

	return string(chars), keywords, len(keywords) > 0
}

func (n *trieNode) FindKeywords(text string) []string {
	chars := []rune(text)
	if len(chars) == 0 {
		return nil
	}

	scopes := n.find(chars)
	return n.collectKeywords(chars, scopes)
}

func (n *trieNode) collectKeywords(chars []rune, scopes []scope) []string {
	set := make(map[string]PlaceholderType)
	for _, v := range scopes {
		set[string(chars[v.start:v.stop])] = Placeholder
	}

	var i int
	keywords := make([]string, len(set))
	for k := range set {
		keywords[i] = k
		i++
	}

	return keywords
}

func (n *trieNode) replaceWithAsterisk(chars []rune, start, stop int) {
	for i := start; i < stop; i++ {
		chars[i] = n.mask
	}
}

// WithMask customizes a Trie with keywords masked as given mask char.
func WithMask(mask rune) TrieOption {
	return func(n *trieNode) {
		n.mask = mask
	}
}

type node struct {
	children map[rune]*node
	fail     *node
	depth    int
	end      bool
}

func (n *node) add(word string) {
	chars := []rune(word)
	if len(chars) == 0 {
		return
	}

	nd := n
	for i, char := range chars {
		if nd.children == nil {
			child := new(node)
			child.depth = i + 1
			nd.children = map[rune]*node{char: child}
			nd = child
		} else if child, ok := nd.children[char]; ok {
			nd = child
		} else {
			child := new(node)
			child.depth = i + 1
			nd.children[char] = child
			nd = child
		}
	}

	nd.end = true
}

func (n *node) build() {
	var nodes []*node
	for _, child := range n.children {
		child.fail = n
		nodes = append(nodes, child)
	}
	for len(nodes) > 0 {
		nd := nodes[0]
		nodes = nodes[1:]
		for key, child := range nd.children {
			nodes = append(nodes, child)
			cur := nd
			for cur != nil {
				if cur.fail == nil {
					child.fail = n
					break
				}
				if fail, ok := cur.fail.children[key]; ok {
					child.fail = fail
					break
				}
				cur = cur.fail
			}
		}
	}
}

func (n *node) find(chars []rune) []scope {
	var scopes []scope
	size := len(chars)
	cur := n

	for i := 0; i < size; i++ {
		child, ok := cur.children[chars[i]]
		if ok {
			cur = child
		} else {
			for cur != n {
				cur = cur.fail
				if child, ok = cur.children[chars[i]]; ok {
					cur = child
					break
				}
			}

			if child == nil {
				continue
			}
		}

		for child != n {
			if child.end {
				scopes = append(scopes, scope{
					start: i + 1 - child.depth,
					stop:  i + 1,
				})
			}
			child = child.fail
		}
	}

	return scopes
}

// Placeholder is a placeholder object that can be used globally.
var Placeholder PlaceholderType

type (
	// AnyType can be used to hold any type.
	AnyType = any
	// PlaceholderType represents a placeholder type.
	PlaceholderType = struct{}
)

// Repr returns the string representation of v.
func Repr(v any) string {
	if v == nil {
		return ""
	}

	// if func (v *Type) String() string, we can't use Elem()
	switch vt := v.(type) {
	case fmt.Stringer:
		return vt.String()
	}

	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	return reprOfValue(val)
}

func reprOfValue(val reflect.Value) string {
	switch vt := val.Interface().(type) {
	case bool:
		return strconv.FormatBool(vt)
	case error:
		return vt.Error()
	case float32:
		return strconv.FormatFloat(float64(vt), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(vt, 'f', -1, 64)
	case fmt.Stringer:
		return vt.String()
	case int:
		return strconv.Itoa(vt)
	case int8:
		return strconv.Itoa(int(vt))
	case int16:
		return strconv.Itoa(int(vt))
	case int32:
		return strconv.Itoa(int(vt))
	case int64:
		return strconv.FormatInt(vt, 10)
	case string:
		return vt
	case uint:
		return strconv.FormatUint(uint64(vt), 10)
	case uint8:
		return strconv.FormatUint(uint64(vt), 10)
	case uint16:
		return strconv.FormatUint(uint64(vt), 10)
	case uint32:
		return strconv.FormatUint(uint64(vt), 10)
	case uint64:
		return strconv.FormatUint(vt, 10)
	case []byte:
		return string(vt)
	default:
		return fmt.Sprint(val.Interface())
	}
}
