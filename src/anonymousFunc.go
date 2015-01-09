package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	intSeqTest1 := intSeq()
	fmt.Println(intSeqTest1())
	fmt.Println(intSeqTest1())
	fmt.Println(intSeqTest1())

	intSeqTest2 := intSeq()
	fmt.Println(intSeqTest2())
}
