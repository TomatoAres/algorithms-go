package main

import "fmt"

func copySlice(s []int) []int {
	//var copyS []int // length and capacity is 0 ，so can not copy
	//var copyS = make([]int,3)
	//var copyS = make([]int,6)
	var copyS = make([]int, 3, 6)
	fmt.Println(copy(copyS, s))
	return copyS
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s, &s)
	fmt.Printf("%v,%p\n", s, &s)

	copyS := copySlice(s)
	fmt.Println(copyS, &copyS)
}
