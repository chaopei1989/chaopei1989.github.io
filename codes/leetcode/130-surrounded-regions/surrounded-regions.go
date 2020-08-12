/*130. 被围绕的区域
题目链接: https://leetcode-cn.com/problems/surrounded-regions/

给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

提示:
图的BFS, DFS
连通性判断-并查集
*/
package main

import "fmt"

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
		if i == 0 || i == len(board)-1 {
			for j := 1; j < len(board[0])-1; j++ {
				dfs(board, i, j)
			}
		}
	}
	// 遍历全部, 将所有 # 换为 O, O 换为 X
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// dfs 遍历所有与 [i,j] 相连的 O
func dfs(board [][]byte, i, j int) {
	// 超出边界
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	// 不是 O, 略过
	if board[i][j] == 'X' {
		return
	}
	// 遍历过了
	if board[i][j] == '#' {
		return
	}
	// 找到连通, 标记表示已遍历过
	board[i][j] = '#'

	dfs(board, i+1, j) // 下
	dfs(board, i, j+1) // 右
	dfs(board, i-1, j) // 上
	dfs(board, i, j-1) // 左
}

/*******************************************************************************
 *
 * 		以下是测试代码
 *
********************************************************************************/

func main() {
	val := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	fmt.Println(val)
	solve(val)
	fmt.Println(val)
}
