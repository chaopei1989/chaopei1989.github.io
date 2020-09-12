package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var res []float64 = make([]float64, 0)
	var queue []*TreeNode = make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		// 当前层节点数
		size := len(queue)
		var average float64
		for i := 0; i < size; i++ {
			top := queue[i]
			average += float64(top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		queue = queue[size:]
		average /= float64(size)
		res = append(res, average)
	}
	return res
}
