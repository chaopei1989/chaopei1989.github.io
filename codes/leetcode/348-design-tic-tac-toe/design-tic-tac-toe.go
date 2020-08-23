package main

type TicTacToe struct {
	n int
	board [][]int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	var ttt = new(TicTacToe)
	ttt.board = make([][]int, n)
	for i := 0; i < n; i++ {
		ttt.board[i] = make([]int, n)
	}
	ttt.n = n
	return *ttt
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	this.board[row][col] = player
	var ret int
	// 判断行
	if ret = this.rowMove(row, col, player); ret != 0 {
		return ret
	}
	// 判断列
	if ret = this.colMove(row, col, player); ret != 0 {
		return ret
	}
	// 判断左上右下对角线
	if ret = this.leftTopMove(row, col, player); ret != 0 {
		return ret
	}
	// 判断右上左下对角线
	if ret = this.leftRightMove(row, col, player); ret != 0 {
		return ret
	}
	return 0
}

func (this *TicTacToe) rowMove(row int, col int, player int) int {
	for i := 0; i < col; i++ {
		if this.board[row][i] != player {
			return 0
		}
	}
	for i := col+1; i < this.n; i++ {
		if this.board[row][i] != player {
			return 0
		}
	}
	return player
}

func (this *TicTacToe) colMove(row int, col int, player int) int {
	for i := 0; i < row; i++ {
		if this.board[i][col] != player {
			return 0
		}
	}
	for i := row+1; i < this.n; i++ {
		if this.board[i][col] != player {
			return 0
		}
	}
	return player
}

func (this *TicTacToe) leftTopMove(row int, col int, player int) int {
	if row == col {
		for i := 0; i < row; i++ {
			if this.board[i][i] != player {
				return 0
			}
		}
		for i := row+1; i < this.n; i++ {
			if this.board[i][i] != player {
				return 0
			}
		}
		return player
	} else {
		return 0
	}
}

func (this *TicTacToe) leftRightMove(row int, col int, player int) int {
	if row + col == this.n-1 {
		for i := 0; i < row; i++ {
			if this.board[i][this.n-1-i] != player {
				return 0
			}
		}
		for i := row+1; i < this.n; i++ {
			if this.board[i][this.n-1-i] != player {
				return 0
			}
		}
		return player
	} else {
		return 0
	}
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */

func main() {

}
