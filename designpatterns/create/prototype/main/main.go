package main

import (
	"fmt"
	"go-learn/designpatterns/create/prototype"
)

func main() {
	t1 := &prototype.Type1{
		Name: "type1",
	}
	t2 := t1.Clone()

	if t1 == t2 {
		fmt.Println("error! getClone not working")

	}

	fmt.Println(t1, "---", &t1)
	fmt.Println(t2, "---", &t2)
}
