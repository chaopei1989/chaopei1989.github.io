package main

import (
	"fmt"
)

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	var matchs1 int
	var matchs1_rev bool
	var matchs2 int
	var lastMatch_s1_i, lastMatch_s1_i_orig = -1, -1
	for s1_i, s2_i := 0, 0; matchs1 < n1 && s1_i < len(s1) && s2_i < len(s2); {
		if s2[s2_i] != s1[s1_i] {
			s1_i++
		} else {
			if lastMatch_s1_i_orig == -1 {
				lastMatch_s1_i_orig = s1_i
			} else if lastMatch_s1_i_orig == s1_i && s2_i == 0 {
				if s1_i > lastMatch_s1_i {
					matchs1_rev = true
					matchs1++
				}
				break
			}
			lastMatch_s1_i = s1_i
			s1_i++
			s2_i++
			// s2 匹配到头了
			if s2_i == len(s2) {
				matchs2++
				s2_i = 0
			}
		}
		if s1_i == len(s1) {
			if s2_i == 0 && matchs2 == 0 {
				// 没有答案
				return 0
			} else {
				matchs1++
				s1_i = 0
			}
		}
	}
	if matchs1 == 0 {
		return 0
	}
	if matchs1_rev {
		return (n1 - 1) * matchs2 / (matchs1 - matchs2)
	} else {
		return n1 * matchs2 / matchs1 / n2
	}
}

func main() {
	// abc abc abc
	// ca
	// 应该是 2
	//fmt.Println(getMaxRepetitions(
	//	"phqghumeaylnlfdxfircvscxggbwkfnqduxwfnfozvsrtkjprepggxrpnrvystmwcysyycqpevikeffmznimkkasvwsrenzkycxf",
	//	100, "xtlsgypsfa", 1))
	fmt.Println(getMaxRepetitions(
		"acb",
		4, "ab", 2))
}
