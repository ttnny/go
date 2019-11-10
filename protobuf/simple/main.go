package main

import (
	"fmt"
	"go/protobuf/simple/pb"
)

func main() {
	fmt.Println("hello world")

	doSimple()
}

func doSimple() {
	sm := pb.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "My Simple Message",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(sm)
	sm.Name = "New Name"

	// Use getters & setters to access the fields
	// because they have safety checks
	fmt.Println(sm.GetId())
}
