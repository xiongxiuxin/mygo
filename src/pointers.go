package main

//指针

import "fmt"

func zeroval(a int) {
	a = 0
}
func zeroptr(a *int) {
	*a = 0
}
func main() {
	a := 1
	fmt.Println("initial:", a)
	zeroval(a)
	fmt.Println("zeroval:", a)
	zeroptr(&a)
	fmt.Println("zeroptr:", a)
	fmt.Println("pointer:", &a)
}
