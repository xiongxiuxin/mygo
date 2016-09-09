package main

import "fmt"

func main() {
	str := "hello,世界"
	ln := len(str)
	for i := 0; i < ln; i++ {
		fmt.Println(i, str[i])
	}
	hello()
}

func hello() {
	str := "hello，世界"
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
