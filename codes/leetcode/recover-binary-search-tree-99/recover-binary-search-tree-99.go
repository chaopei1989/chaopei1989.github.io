/*99. 恢复二叉搜索树
题目链接: https://leetcode-cn.com/problems/recover-binary-search-tree/

二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？

提示: 先考虑中序遍历
1324 -> 123
14325 -> 12345
21 -> 12
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*******************************************************************************
 *
 * 		递归中序遍历
 *
********************************************************************************/

func recoverTreeWithParent(env *Env, root *TreeNode, parent *TreeNode) bool {
	// 1. 中序遍历
	if root.Left != nil {
		if recoverTreeWithParent(env, root.Left, root) {
			return true
		}
	}
	var ret = trace(env, root, parent)
	// 找到了
	if ret {
		doReverse(env)
		return true
	}
	if root.Right != nil {
		if recoverTreeWithParent(env, root.Right, root) {
			return true
		}
	}
	return false
}

func doReverse(env *Env) {
	if env.keyNodeL == nil || env.keyNodeR == nil {
		return
	}
	// 交换节点
	env.keyNodeL.node.Val ^= env.keyNodeR.node.Val
	env.keyNodeR.node.Val ^= env.keyNodeL.node.Val
	env.keyNodeL.node.Val ^= env.keyNodeR.node.Val
}

func recoverTree(root *TreeNode) {
	var env = new(Env)
	if !recoverTreeWithParent(env, root, nil) {
		env.keyNodeR = env.keyNodeLRight
		doReverse(env)
	}
}

type ProblemNode struct {
	node *TreeNode
}

type Env struct {
	last          *TreeNode
	keyNodeL      *ProblemNode
	keyNodeLRight *ProblemNode
	keyNodeR      *ProblemNode
}

func trace(env *Env, node *TreeNode, parent *TreeNode) bool {
	if env.last != nil {
		if env.last.Val > node.Val {
			if env.keyNodeL == nil {
				env.keyNodeL = new(ProblemNode)
				env.keyNodeL.node = env.last
				env.keyNodeLRight = new(ProblemNode)
				env.keyNodeLRight.node = node
			} else {
				env.keyNodeR = new(ProblemNode)
				env.keyNodeR.node = node
			}
		}
		if env.keyNodeL != nil && env.keyNodeR != nil {
			return true
		}
	}
	// next
	env.last = node

	return false
}

/*******************************************************************************
 *
 * 		下面是测试代码
 *
********************************************************************************/

func main() {
	// [1,3,null,null,2]
	//var root = new(TreeNode)
	//root.Val = 1
	//var rL = new(TreeNode)
	//rL.Val = 3
	//var rLR = new(TreeNode)
	//rLR.Val = 2
	//root.Left = rL
	//rL.Right = rLR

	// [3,1,4,null,null,2]
	var root = new(TreeNode)
	root.Val = 3
	var rL = new(TreeNode)
	rL.Val = 1
	var rR = new(TreeNode)
	rR.Val = 4
	var rRL = new(TreeNode)
	rRL.Val = 2
	root.Left = rL
	root.Right = rR
	rR.Left = rRL

	printTree(root)
	fmt.Println()
	recoverTree(root)
	printTree(root)
}

// BFS 打印结果
func printTreeBFS(queue []*TreeNode, lastNode *TreeNode) {
	if len(queue) > 0 {
		tmp := queue[0]
		if tmp != nil {
			fmt.Printf("%d ", tmp.Val)
			queue = append(queue, tmp.Left)
			if tmp.Left != nil {
				lastNode = tmp.Left
			}
			queue = append(queue, tmp.Right)
			if tmp.Right != nil {
				lastNode = tmp.Right
			}
		} else {
			fmt.Printf("nil ")
			queue = append(queue, nil)
			queue = append(queue, nil)
		}
		queue = queue[1:]
		if tmp != lastNode {
			printTreeBFS(queue, lastNode)
		}
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTreeBFS([]*TreeNode{root}, root)
}
