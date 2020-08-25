/*394. 字符串解码
题目链接: https://leetcode-cn.com/problems/decode-string/

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/
package main

import "fmt"

// TokenType Token的类型
type TokenType int32

const (
	// TypeInvalid 默认无效类型
	TypeInvalid TokenType = -1
	// TypeAbc 字母序列
	TypeAbc TokenType = 0
	// TypeInt 数字序列
	TypeInt TokenType = 1
	// TypeLeftBracket 左括号
	TypeLeftBracket TokenType = 2
	// TypeRightBracket 左括号
	TypeRightBracket TokenType = 3
)

// Token 词法分析的Token
type Token struct {
	// 原始字符串
	src string
	// 解析中间结果, 是个String的列表
	tokenType TokenType
}

// 按照单行解析
func lexical(res *[]*Token, line string) {
	// tokens := make([]*Token, 0)
	predict := TypeInvalid
	if len(line) > 0 {
		// 字符类型
		predict = charType(line[0])
	}
	var tmp *Token
	var rest string
	for i := 1; i < len(line); i++ {
		// 一直重复到类型不同, 添加token到结果
		if charType(line[i]) != predict {
			// i 所代表的的不是想要的
			tmp = new(Token)
			tmp.src = line[0:i]
			tmp.tokenType = predict
			rest = line[i:len(line)]
			break
		} else {
			if i == len(line)-1 {
				// 最后一个字符
				tmp = new(Token)
				tmp.src = line
				tmp.tokenType = predict
				break
			}
			// else 积累, 下一个字符
		}
	}
	*res = append(*res, tmp)
	if len(rest) > 0 {
		lexical(res, rest)
	}
}

func charType(ch byte) TokenType {
	var predict TokenType
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		predict = TypeAbc
	} else if ch >= '0' && ch <= '9' {
		predict = TypeInt
	} else if ch == '[' {
		predict = TypeLeftBracket
	} else if ch == '[' {
		predict = TypeLeftBracket
	} else if ch == ']' {
		predict = TypeRightBracket
	} else {
		predict = TypeInvalid
	}
	return predict
}

// 核心函数
func decodeString(s string) string {
	// 格式总是对的, 一共两种形式, 一种是 3[ab], 一种是 cd
	// 逐字读取, 状态为结束时出token list
	tokens := make([]*Token, 0)
	lexical(&tokens, s)
	for _, t := range tokens {
		fmt.Printf("%s, %d\n", t.src, t.tokenType)
	}
	// 构建语法树

	// 解析语法树, 算结果
	return ""
}

// 测试入口
func main() {
	decodeString("2[abc]3[cd]ef")
}
