package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)

	// pop top, visit
	// 每次都是遍历当前level的, queue中当前存的就是同一level的节点
	for len(queue) > 0 {
		lvlSize := len(queue)
		lvlRes := make([]int, 0)
		for i := 0; i < lvlSize; i++ {
			top := queue[0]
			lvlRes = append(lvlRes, top.Val)
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		resNew := make([][]int, 0)
		resNew = append(resNew, lvlRes)
		res = append(resNew, res...)
	}
	return res
}


func main() {

}
