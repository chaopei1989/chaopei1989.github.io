package main

import "fmt"

func solveNQueens(n int) [][]string {
	columns := make(map[int]bool, 0)
	diag1 := make(map[int]bool, 0)
	diag2 := make(map[int]bool, 0)
	res := solveNQueensRecursive(columns, diag1, diag2, n, 0, nil, nil)
	return res
}

func generateBoard(queens [][]int, n int) []string {
	board := make([]string, n)
	for i := 0; i < len(queens); i++ {
		// queens[i][0] 行
		// queens[i][1] 列
		tmp := make([]byte, n)
		for j := 0; j < n; j++ {
			if j != queens[i][1] {
				tmp[j] = '.'
			} else {
				tmp[j] = 'Q'
			}
		}
		board[queens[i][0]] = string(tmp)
	}
	return board
}

func solveNQueensRecursive(columns, diag1, diag2 map[int]bool, n int, row int, curr [][]int, res [][]string) [][]string {
	// 第 row 行
	if row == n && len(curr) == n {
		fmt.Println(curr)
		return append(res, generateBoard(curr, n))
	}
	// 遍历列
	for i := 0; i < n; i++ {
		if v, ok := columns[i]; ok && v {
			continue
		}
		if v1, ok1 := diag1[row+i]; ok1 && v1 {
			continue
		}
		if v2, ok2 := diag2[row-i]; ok2 && v2 {
			continue
		}

		columns[i] = true
		diag1[row+i] = true
		diag2[row-i] = true
		curr = append(curr, []int{row, i})
		res = solveNQueensRecursive(columns, diag1, diag2, n, row+1, curr, res)
		curr = curr[0 : len(curr)-1]
		columns[i] = false
		diag1[row+i] = false
		diag2[row-i] = false
	}
	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
