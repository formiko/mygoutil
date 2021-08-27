package mygoutil

import (
	"sort"
)

// 对字符串本身的字符序列排序
func SortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}