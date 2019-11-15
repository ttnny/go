package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go/protobuf/simple/pb"
	"io/ioutil"
	"log"
)

func main() {
	sm1 := doSimple()
	writeToFile("test.bin", sm1)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("test.bin", sm2)
	fmt.Println(sm2)

	smJSON := toJSON(sm2)
	fmt.Println(smJSON)

	fromJSON(smJSON, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protobuf struct", err)
		return err2
	}

	return nil
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}

	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON")
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal JSON into the pb struct", err)
	}
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(sm)
	sm.Name = "New Name"

	// Use getters & setters to access the fields
	// because they have safety checks
	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
