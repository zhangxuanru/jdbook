/*
@Time : 2020/4/15 15:28
@Author : zxr
@File : 5.3.1
@Software: GoLand
*/
package string

import "unicode"

//如何实现单词反转
func ReverseWorld(str string) string {
	r := []rune(str)
	begin := 0
	end := len(r) - 1
	swap(r, 0, end)
	for k, v := range r {
		if unicode.IsSpace(v) {
			swap(r, begin, k-1)
			begin = k + 1
		}
	}
	swap(r, begin, end)
	return string(r)
}

func swap(r []rune, begin, end int) {
	for begin < end {
		r[begin], r[end] = r[end], r[begin]
		begin++
		end--
	}
}
