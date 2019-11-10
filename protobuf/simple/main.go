package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go/protobuf/simple/pb"
	"io/ioutil"
	"log"
)

func main() {
	sm := doSimple()

	writeToFile()

	readFromFile()
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if (err != nil) {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
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

	return &sm
}
