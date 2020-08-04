/*114. 二叉树展开为链表
题目链接: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

给定一个二叉树，原地将它展开为一个单链表。

例如，给定二叉树
    1
   / \
  2   5
 / \   \
3   4   6

将其展开为：
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/
package main

// TreeNode represents as a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			flatten(root.Left)
		}
		if root.Right != nil {
			flatten(root.Right)
		}
		tmpOrigRight := root.Right
		root.Right = root.Left
		root.Left = nil

		tmp := root
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = tmpOrigRight
	} else {
		// do nothing
	}
}

func main() {
	// do nothing
}
