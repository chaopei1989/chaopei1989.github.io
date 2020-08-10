/*696. 计数二进制子串
题目链接: https://leetcode-cn.com/problems/count-binary-substrings/

给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，
并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。

示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
示例 2 :

输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
注意：

s.length 在1到50,000之间。
s 只包含“0”或“1”字符。

提示: 将0和1分别计数存数组
*/
package main

import "fmt"

func countBinarySubstrings(s string) int {
	var counts []int = make([]int, 0)
	// 1. 先遍历, 计数0和1
	var count int = 0
	var ch byte = 0
	for i := 0; i < len(s); i++ {
		if ch == s[i] {
			count++
		} else if ch != 0 {
			counts = append(counts, count)
			ch = s[i]
			count = 1
		} else {
			ch = s[i]
			count = 1
		}
	}
	counts = append(counts, count)
	// 2. 遍历counts数组, 计算相邻数字的最小值
	var ret int = 0
	for i := 0; i < len(counts)-1; i++ {
		if counts[i] < counts[i+1] {
			ret += counts[i]
		} else {
			ret += counts[i+1]
		}
	}
	return ret
}

func main() {
	ret := countBinarySubstrings("10101")
	fmt.Println(ret)
}
