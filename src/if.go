package main

import "fmt"

func main() {
	var x = iftest(5)
	fmt.Println(x)
}
func iftest(x int) int {
	if x < 10 {
		return x
	} else {
		return 10
	}
}
