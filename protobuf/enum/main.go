package main

import (
	"fmt"
	enumpb "go/protobuf/enum/pb"
)

func main() {
	doEnum()
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}

	fmt.Println(em)
}