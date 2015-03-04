package main

import (
	"fmt"
	"time"
)

func work(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go work(done)
	fmt.Println("main")
	if <-done {
		fmt.Println("ok")
	} else {
		fmt.Println("failed")
	}
	fmt.Println("end")
}
