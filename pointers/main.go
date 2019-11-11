package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-------- ARRAYS")
	arrays()

	fmt.Println("-------- SLICES")
	slices()
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("array nums	: %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("incr nums	: %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incr nums	: %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}

func slices() {
	dirs := []string{"up", "right", "bottom", "left"}

	upperCase(dirs)
	fmt.Printf("original list	: %p %q\n", &dirs, dirs)

	upperCasePtr(&dirs)
	fmt.Printf("original list	: %p %q\n", &dirs, dirs)
}

func upperCase(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	list = append(list, "CENTER")

	fmt.Printf("upperCase list	: %p %q\n", &list, list)
}

func upperCasePtr(list *[]string) {
	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "CENTER")

	fmt.Printf("upperCase list	: %p %q\n", list, list)
}