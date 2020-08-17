package dfa

import (
	"strings"
)

type DFA interface {
	Add(words ...string)
	Filter(word string) string
	AddSkip(words... string)
}

type node struct {
	IsEnd bool
	value map[string]node
}

func (n *node) hasKey(key string) bool {
	_, ok := n.value[key]
	return ok
}

type dfa struct {
	skipRoot []string
	root     *node
}

func (d dfa) inSkipRoot(key string) bool {
	for _, v:= range d.skipRoot{
		if v == key {
			return true
		}
	}
	return false
}

func (d *dfa) Add(words ...string) {
	for _, word := range words {
		nowNode := d.root
		wordLength := len([]rune(word))
		for i, w := range []rune(word) {
			key := string(w)
			if d.root.hasKey(key) {
				*nowNode = d.root.value[key]
				nowNode.IsEnd = false
			} else {
				newNode := node{
					IsEnd: false,
					value: make(map[string]node, 0),
				}
				if wordLength == i+1 {
					newNode.IsEnd = true
				} else {
					d.root.IsEnd = false
				}
				nowNode.value[key] = newNode
				nowNode = &newNode
			}
		}
	}
}

func (d dfa) check(beginIndex int, word string) (matchedIndex int) {
	n := *d.root
	flag  := false
	for _, w := range []rune(word)[beginIndex:] {
		key := string(w)
		if d.inSkipRoot(key) { //防止特殊字符干扰
			matchedIndex += 1
			continue
		}
		if n.hasKey(key) {
			n = n.value[key]
			matchedIndex += 1
			if n.IsEnd {
				flag = true
				break
			}
		} else {
			break
		}
	}
	if matchedIndex < 2 || !flag{
		return 0
	}
	return matchedIndex
}

func (d dfa) get(word string) []string {
	matchedList := make([]string, 0)
	for i, _ := range []rune(word) {
		matched := d.check(i, word)
		if matched > 0 {
			matchedWord := []rune(word)[i:i+matched]
			matchedList = append(matchedList, string(matchedWord))
		}
	}
	return matchedList
}

func (d dfa) Filter(word string) string {
	matchedWordList := d.get(word)

	if len(matchedWordList) >  0 {
		for _, m := range matchedWordList {
			replaceStr := ""
			for _, s := range []rune(m) {
				_ = s
				replaceStr += "*"
			}
			word =  strings.Replace(word, m, replaceStr, 1)
		}
		return word
	}
	return word
}

func (d *dfa)AddSkip(words... string)  {
	d.skipRoot = append(d.skipRoot,words...)
}

func NewDFA() DFA {
	return &dfa{
		skipRoot: make([]string,0),
		root: &node{
			IsEnd: false,
			value: make(map[string]node, 0),
		},
	}
}
