```
20. 有效的括号

题目链接: https://leetcode-cn.com/problems/valid-parentheses/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
```

## 提示

1. 使用入栈来存储左括号, 出栈来匹配右括号.
2. 使用 Map 存储括号的匹配关系.
3. 从左往右依次遍历字符串, 当匹配错误时即可退出程序.
4. 一开始可以判断总长度为奇数时, 即匹配错误.

## 做法一

