package main

import "fmt"

func main() {
	arr := [10]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	for i, v := range arr {
		fmt.Println("arr[", i, "]=", v)
	}
	modify(arr)
	fmt.Println("main arr is :", arr)
	var twoDB [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDB[i][j] = i + j
		}
	}
	fmt.Println(twoDB)
}
func modify(arr [10]int) {
	arr[0] = 99
	fmt.Println("modify arr is :", arr)
}
