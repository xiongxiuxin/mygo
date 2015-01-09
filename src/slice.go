package main

import "fmt"

func main() {
	var myArr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var mySlice1 []int = myArr[:5]
	mySlice2 := myArr[5:]
	//fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	var mySlice3 = make([]int, 5)
	fmt.Println("mySlice3 lenth is:", len(mySlice3), "mySlice3 capacity is:", cap(mySlice3))
	fmt.Println(mySlice3)
	mySlice4 := make([]int, 5, 10)
	fmt.Println("mySlice4 lenth is:", len(mySlice4), "mySlice4 capacity is:", cap(mySlice4))
	mySlice5 := []int{1, 2, 3, 4, 5}
	mySlice1 = append(mySlice1, mySlice5...)
	fmt.Println(mySlice1)
	fmt.Println(mySlice5)
	for i := 0; i < len(mySlice5); i++ {
		fmt.Println(mySlice5[i])
	}
	for i, v := range mySlice5 {
		fmt.Println(i, v)
	}

	oldSlice := make([]int, 5, 10)
	for i := 0; i < len(oldSlice); i++ {
		oldSlice[i] = i
	}
	newSlice := oldSlice[:10]
	fmt.Println(newSlice)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	//copy(slice1, slice2)
	//fmt.Println(slice1, "-----", slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, "-----", slice2)

	twoDB := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDB[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDB[i][j] = i + j
		}
	}
	fmt.Println(twoDB)
}
