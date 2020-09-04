package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePathsRecursive(node, other *TreeNode, curr string, res *[]string) {
	if node == nil {
		if other == nil {
			*res = append(*res, curr)
		}
		return
	}
	if curr != "" {
		curr = curr + "->"
	}
	curr = curr + strconv.Itoa(node.Val)
	if node.Left == nil {
		binaryTreePathsRecursive(node.Right, node.Left, curr, res)
	} else if node.Right == nil {
		binaryTreePathsRecursive(node.Left, node.Right, curr, res)
	} else {
		binaryTreePathsRecursive(node.Left, node.Right, curr, res)
		binaryTreePathsRecursive(node.Right, node.Left, curr, res)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := make([]string, 0)
	binaryTreePathsRecursive(root, nil, "", &res)
	return res
}

func main() {

}
