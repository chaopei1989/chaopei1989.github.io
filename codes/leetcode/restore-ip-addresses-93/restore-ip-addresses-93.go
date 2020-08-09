/*93. 复原IP地址
题目链接: https://leetcode-cn.com/problems/restore-ip-addresses/

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。



示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var res []string = make([]string, 4)

func restoreIpAddresses(s string, index int, size int, result []string) {
	if size > 3 || index+size > len(s) {
		// oob
		return
	}
	if s[index] == '0' && size > 1 {
		// 0 alone
		return
	}
	if s[index] > '2' && size > 2 {
		// 3X
		return
	}
	d, _ := strconv.Atoi(s[index : index+size])
	if d > 255 {
		return
	}
	res = append(res, s[index:index+size])
	if len(res) < 4 {
		for i := 1; i <= len(s)-index-size && i <= 3; i++ {
			restoreIpAddresses(s, index+size, i, result)
		}
	} else {
		if index == len(s) {
			// TODO output
			tmp := strings.Join(res, ".")
			result = append(result, tmp)
		}
	}
	res = res[0 : len(res)-1]
}

func main() {
	var r = "25525511135"
	var result []string = make([]string, 0)
	restoreIpAddresses(r, 0, 1, result)
	fmt.Println(result)
}
