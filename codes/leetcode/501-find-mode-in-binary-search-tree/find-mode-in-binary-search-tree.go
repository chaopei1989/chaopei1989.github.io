package find_mode_in_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	Val       []int
	Count     int
	CurrVal   int
	CurrCount int
}

func findModeRec(root *TreeNode, res *Result) {
	// 中序遍历
	if root.Left != nil {
		findModeRec(root.Left, res)
	}
	res.update(root)
	if root.Right != nil {
		findModeRec(root.Right, res)
	}
}
func (res *Result) update(node *TreeNode) {
	if res.CurrCount > 0 && res.CurrVal == node.Val {
		res.CurrCount++
	} else {
		res.CurrVal = node.Val
		res.CurrCount = 1
	}
	if res.Count == 0 || res.CurrCount == res.Count {
		res.Val = append(res.Val, res.CurrVal)
		res.Count = res.CurrCount
	} else if res.CurrCount > res.Count {
		res.Val = []int{res.CurrVal}
		res.Count = res.CurrCount
	} else {
		// ignore
	}
}

func findMode(root *TreeNode) []int {
	res := new(Result)
	findModeRec(root, res)
	return res.Val
}

func main() {

}
