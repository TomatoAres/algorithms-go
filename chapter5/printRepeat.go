package main

import (
	"fmt"
	"sort"
)

func printRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		fmt.Println("i=", i, " value=", data[i])
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(" ", data[i])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	fmt.Println("data:", data)
	fmt.Print("Repeating elements are : ")
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(" ", data[i])
		}
	}
	fmt.Println()
}

//=========================

type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}
func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}
func printRepeating3(data []int) {
	s := make(Set)
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if s.Find(data[i]) {
			fmt.Print(" ", data[i])
		} else {
			s.Add(data[i])
		}
	}
	fmt.Println()
}

func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	// 书中有误
	for i := 0; i < intrange; i++ {
		count[i] = 0
	}
	fmt.Println("count:",count)

	fmt.Println("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}

func main() {
	data := []int{1, 2, 3, 1, 1, 2, 2, 3, 1}
	printRepeating4(data, 4)
}
