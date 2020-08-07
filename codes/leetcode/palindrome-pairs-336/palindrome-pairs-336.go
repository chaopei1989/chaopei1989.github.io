/*336. 回文对
题目链接: https://leetcode-cn.com/problems/palindrome-pairs/

给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。

示例 1:

输入: ["abcd","dcba","lls","s","sssll"]
输出: [[0,1],[1,0],[3,2],[2,4]]
解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
示例 2:

输入: ["bat","tab","cat"]
输出: [[0,1],[1,0]]
解释: 可拼接成的回文串为 ["battab","tabbat"]
*/
package main

import "fmt"

type Node struct {
	// 记录该节点对应的字母
	v     byte
	index int
	// 记录子节点
	children [26]*Node
}

// 指定根节点, 构建字典树, 使用树形结构
func buildTrie(root *Node, words []string) {
	// 遍历 words
	for index, word := range words {
		var len = len(word)
		// 允许有空字符串
		if len == 0 {
			root.index = index
		}
		var curr = root
		for i := 0; i < len; i++ {
			var x = int(word[i] - 'a')
			var node = curr.children[x]
			if node == nil {
				node = new(Node)
				node.index = -1
				node.v = word[i]
				curr.children[x] = node
			}
			if i == len-1 {
				node.index = index
			}
			curr = node
		}
	}
}

// 判断是否回文 [left, right]
func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func hasWord(root *Node, s string, left, right int, rev bool) int {
	var curr = root
	if left <= right {
		if !rev {
			for i := left; i <= right; i++ {
				curr = curr.children[int(s[i]-'a')]
				if curr == nil {
					break
				}
			}
		} else {
			for i := right; i >= left; i-- {
				curr = curr.children[int(s[i]-'a')]
				if curr == nil {
					break
				}
			}
		}
	}
	if curr != nil && curr.index >= 0 {
		return curr.index
	} else {
		return -1
	}
}

func palindromePairs(words []string) [][]int {
	var res = make([][]int, 0)
	// 1. 构建字典树
	var root *Node
	root = new(Node)
	root.index = -1
	buildTrie(root, words)

	// 2. 遍历每一个词, 计算其是否存在回文的前缀和后缀
	for index, word := range words {
		for i := 0; i <= len(word); i++ {
			// 前缀回文, 存在空串
			if i != 0 && isPalindrome(word, 0, i-1) {
				// 查找后缀的反转是否存在
				var hw = hasWord(root, word, i, len(word)-1, true)
				if hw >= 0 && hw != index {
					res = append(res, []int{hw, index})
				}
			}
			// 后缀回文, 存在空串
			if isPalindrome(word, i, len(word)-1) {
				// 查找前缀的反转是否存在
				var hw = hasWord(root, word, 0, i-1, true)
				if hw >= 0 && hw != index {
					res = append(res, []int{index, hw})
				}
			}
		}
	}

	return res
}

func main() {
	var res [][]int
	res = palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"})
	fmt.Println(res)
	res = palindromePairs([]string{"aba", "ba", "a", "caba"})
	fmt.Println(res)
}
