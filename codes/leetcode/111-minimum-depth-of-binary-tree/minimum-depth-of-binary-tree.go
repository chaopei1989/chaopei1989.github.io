package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left != 0 && right != 0 {
		// 对比左右节点更小的
		if left < right {
			return left + 1
		} else {
			return right + 1
		}
	} else if left == 0 && right == 0 {
		// 左右节点均为空
		return 1
	} else {
		// 左右节点有一方为空, 意味着当前节点不能算作是叶子节点.
		if left < right {
			return right + 1
		} else {
			return left + 1
		}
	}
}

func main()  {
	
}