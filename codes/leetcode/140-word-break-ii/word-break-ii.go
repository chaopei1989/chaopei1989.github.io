package main

import "strings"

func wordBreak(s string, wordDict []string) []string {
	// 记录从index(key)到最后的结果(value)
	rec := make(map[int][][]string)
	// 词库
	wordset := make(map[string]bool)
	for _, word := range wordDict {
		wordset[word] = true
	}
	res := backtrack(wordset, rec, s, 0)
	ret := make([]string, 0)
	for _, words := range res {
		ret = append(ret, strings.Join(words, " "))
	}
	return ret
}


func backtrack(wordset map[string]bool, rec map[int][][]string, s string, index int) [][]string {
	if rec[index] != nil {
		return rec[index]
	}
	res := make([][]string, 0)
	for i := index+1; i < len(s); i++ {
		// 从index到i的子串
		substr := s[index:i]
		if v, ok := wordset[substr]; ok && v {
			for _, nextWords := range backtrack(wordset, rec, s, i) {
				res = append(res, append([]string{substr}, nextWords...))
			}
		}
	}
	// 最后一个字母
	word := s[index:]
	if v, ok := wordset[word]; ok && v {
		res = append(res, []string{word})
	}
	rec[index] = res
	return res
}

func main() {

}
