package main

type PoiMap map[int]bool

func (self PoiMap) exist(i, j int) bool {
	if v, ok := self[i*1000+j]; ok && v {
		return true
	}
	return false
}

func (self PoiMap) insert(i, j int) {
	self[i*1000+j] = true
}

func (self PoiMap) del(i, j int) {
	self[i*1000+j] = false
}

func existDFS(c, i, j int, trace PoiMap, board [][]byte, word string) bool {
	if word[c] != board[i][j] {
		return false
	}
	trace.insert(i, j)
	if c+1 == len(word) {
		return true
	}
	// 上
	if i-1 >= 0 && !trace.exist(i-1, j) {
		if existDFS(c+1, i-1, j, trace, board, word) {
			return true
		}
	}
	// 下
	if i+1 < len(board) && !trace.exist(i+1, j) {
		if existDFS(c+1, i+1, j, trace, board, word) {
			return true
		}
	}
	// 左
	if j-1 >= 0 && !trace.exist(i, j-1) {
		if existDFS(c+1, i, j-1, trace, board, word) {
			return true
		}
	}
	// 右
	if j+1 < len(board[0]) && !trace.exist(i, j+1) {
		if existDFS(c+1, i, j+1, trace, board, word) {
			return true
		}
	}
	trace.del(i, j)
	return false
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if existDFS(0, i, j, make(PoiMap, 0), board, word) {
					return true
				}
			}
		}
	}
	return false
}

func main() {

}
