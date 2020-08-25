package main

import (
	"fmt"
	"os"
	"runtime"
)

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

func f(x int) {
	//if x == 0 {
	//	return
	//}
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0, defer will finish
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	fmt.Printf("main: %d\n", double(4))
	c1 := make(chan int)
	func(ch chan int) {
		fmt.Println("main, inner func", c1 == ch)
	}(c1)
	defer printStack()
	f(3)
}
