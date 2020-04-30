/*
@Time : 2020/4/9 17:46
@Author : zxr
@File : 5.1
@Software: GoLand
*/
package string

import (
	"fmt"

	"github.com/isdamir/gotype"
)

//如何求一个字符串的所有排列
//https://leetcode-cn.com/problems/permutation-i-lcci/

//error
func PermutationStr(str string) []string {
	r := make([]string, 1)
	r[0] = str
	n := len(str)
	sRune := []rune(str)
	start := 1
	for start <= n {
		for i := 1; i < n-1; i++ {
			sRune[i], sRune[i+1] = sRune[i+1], sRune[i]
			r = append(r, string(sRune))
		}
		if start == n {
			break
		}
		sRune = []rune(str)
		sRune[start], sRune[0] = sRune[0], sRune[start]
		r = append(r, string(sRune))
		start++
	}
	return r
}

//success
func Permutation2(S string) []string {
	n := len(S)
	ret := []string{}
	visited := map[int]bool{}
	var br func(now string, i int)

	br = func(now string, k int) {
		fmt.Println("params:", now, "i=", k)
		if len(now) == n {
			ret = append(ret, now)
			fmt.Println("now visited:", visited, "i=", k)
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			br(now+string(S[i]), i)
			visited[i] = false
			fmt.Println("i=", i, "visited:", visited)
		}
	}
	br("", 0)
	return ret
}

//success
func Permutation3(str []rune, start int) {
	if str == nil {
		return
	}
	if start == len(str)-1 {
		fmt.Println(string(str))
	} else {
		for i := start; i < len(str); i++ {
			gotype.SwapRune(str, start, i)
			Permutation3(str, start+1)
			gotype.SwapRune(str, start, i)
		}
	}
}
