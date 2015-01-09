package main

import (
	"fmt"
)

type PersonInfo struct {
	ID     string
	Name   string
	Mobile string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo, 10)
	personDB["1"] = PersonInfo{"1", "xxx", "13879787878"}
	personDB["2"] = PersonInfo{"2", "zzz", "13879787868"}
	person, ok := personDB["2"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 2.")
		delete(personDB, "2")
	} else {
		fmt.Println("Did not find person with ID 2")
	}
	fmt.Println(personDB)
}
