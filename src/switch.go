package main

import (
	"fmt"
	"time"
)

func main() {
	//test(8)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is weekday")
	}
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Now is before noon")
	default:
		fmt.Println("Now is after noon")
	}
}

func test(x int) {
	switch x {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4,5,6")
	default:
		fmt.Println("Default")
	}
}
