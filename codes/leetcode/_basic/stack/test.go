package main

import "fmt"

func double(x int) (result int) {
	defer func() {
		result++
		fmt.Printf("double 1(%d) = %d\n", x, result)
	}()
	defer func() {
		result++
		fmt.Printf("double 2(%d) = %d\n", x, result)
	}()
	fmt.Printf("double 3: %d\n", result)
	return x + x
}

func main() {
	fmt.Printf("main: %d\n", double(4))
}
