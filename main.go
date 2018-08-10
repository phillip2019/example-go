package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// fmt.Print("I am happy!")
	// defer printStack()
	// f(18)
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/(x-16))
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

// 欧几里得算法
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波那契数列(Fibonacci)第N个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
