/*43. 字符串相乘
题目链接: https://leetcode-cn.com/problems/multiply-strings/

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(num1 string, num2 string) string {
	var res string = ""
	var jw bool = false
	for i := 0; i < len(num1) || i < len(num2); i++ {
		var c1, c2 int
		if i < len(num1) {
			c1 = int(num1[len(num1)-1-i] - '0')
		} else {
			c1 = 0
		}
		if i < len(num2) {
			c2 = int(num2[len(num2)-1-i] - '0')
		} else {
			c2 = 0
		}
		var curr = c1 + c2
		if jw {
			curr++
		}
		var currStr string = strconv.Itoa(curr)
		if curr > 9 {
			jw = true
			res = currStr[1:] + res
		} else {
			jw = false
			res = currStr + res
		}
	}
	if jw {
		res = "1" + res
	}
	return res
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var res string = ""
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			var tmp = strconv.Itoa(int(num1[i]-'0') * int(num2[j]-'0'))
			tmp += strings.Repeat("0", len(num1)-i-1+len(num2)-j-1)
			res = add(res, tmp)
		}
	}
	return res
}

func main() {
	fmt.Println(multiply("8764", "454"))
}
