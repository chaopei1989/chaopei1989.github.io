package main

import "fmt"

var (
	visited []bool
)

func visit(rooms [][]int, visited []bool, curr int)  {
	if !visited[curr] {
		visited[curr] = true
		for _, r := range rooms[curr] {
			visit(rooms, visited, r)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	visited = make([]bool, len(rooms))
	visit(rooms, visited, 0)
	for _, r := range visited {
		if !r {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{},
	}))
}
