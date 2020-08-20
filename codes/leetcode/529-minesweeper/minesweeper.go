package main

// dfs 遍历图
func updateBoard(board [][]byte, click []int) [][]byte {
	x, y:= click[0], click[1]
	// 超过了边界
	if x < 0 || x >= len(board) {
		return board
	}
	if y < 0 || y >= len(board[0]) {
		return board
	}
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	// 遍历过了
	if board[x][y] != 'E' {
		return board
	}
	
	var mine byte = '0'
	// 上
	if x-1 >= 0 && board[x-1][y] == 'M' {
		mine++
	}
	// 下
	if x+1 < len(board) && board[x+1][y] == 'M' {
		mine++
	}
	// 左
	if y-1 >= 0 && board[x][y-1] == 'M' {
		mine++
	}
	// 右
	if y+1 < len(board[x]) && board[x][y+1] == 'M' {
		mine++
	}

	// 左上
	if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] == 'M' {
		mine++
	}
	// 右上
	if x-1 >= 0 && y+1 < len(board[x]) && board[x-1][y+1] == 'M' {
		mine++
	}
	// 左下
	if x+1 < len(board) && y-1 >= 0 && board[x+1][y-1] == 'M' {
		mine++
	}
	// 右下
	if x+1 < len(board) && y+1 < len(board[x]) && board[x+1][y+1] == 'M' {
		mine++
	}
	if mine > '0' {
		board[x][y] = mine
	} else {
		board[x][y] = 'B'
		board = updateBoard(board, []int{x-1,y})
		board = updateBoard(board, []int{x+1,y})
		board = updateBoard(board, []int{x,y-1})
		board = updateBoard(board, []int{x,y+1})

		board = updateBoard(board, []int{x-1,y-1})
		board = updateBoard(board, []int{x+1,y-1})
		board = updateBoard(board, []int{x-1,y+1})
		board = updateBoard(board, []int{x+1,y+1})
	}
	return board
}

func main() {

}
