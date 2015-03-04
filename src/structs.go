package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"xxx", 10})
	fmt.Println(person{name: "zzz", age: 20})
	fmt.Println(person{name: "aaa"})
	fmt.Println(&person{name: "vvv", age: 30})
	s := person{name: "bbb", age: 40}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 50
	fmt.Println(sp.age, s.age)
}
