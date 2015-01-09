package main

import (
	"fmt"
	"otherPackage"
)

func main() {
	_, b, sum := test1(1, 2)
	fmt.Println(b, sum)
	fmt.Println(otherPackage.fact(4))
	myfunc()
}

//多返回值
func test1(a int, b int) (int, int, int) {
	return b, a, a + b
}

//递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//goto
func myfunc() {
	var a = 0
HERE:
	fmt.Println(a)
	a++
	if a < 10 {
		goto HERE
	}
}
