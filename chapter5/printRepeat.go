package main

import (
	"fmt"
	"sort"
)
var data = []int{1, 2, 3, 1, 1, 2, 2, 3, 1}

// i 0:1 1 1
// 	 1: 2 2
//   2: 3
//   3: 1 1
//	 4: 1
//	 5:2
//	 6:
//   7:
// out:1 1 1 2 2 3 1 1 1 2
// 输出冗余，可能不符合需求
func printRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		//fmt.Printf("data[%d]=%d:",i,data[i])
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {

				fmt.Print(" ", data[i])
			}
		}
		//fmt.Println()
	}
	fmt.Println()
}

// 先排序，再取重
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

//hash 键唯一的特性
func printRepeating3(data []int) {
	s := make(Set)
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		//fmt.Println(s[data[i]]) // 用于找规律
		// 判断方法1： Find 说明之前存在，重复了
		//if s.Find(data[i]) {
		// 方法2： 未重复 false，重复 true
		if s[data[i]] {
			fmt.Print(" ", data[i])
		} else {
			s.Add(data[i])
		}
	}
	fmt.Println()

	//fmt.Println("printRepeating3:",s)	// debug
}

// intrange counting 须知的范围
//counting：需知所有值所在的范围 例 a中所有的范围就是 1-3，数组下标最大为3，intrange=4
// 其实和hash一样，本质上用列表实现了hash
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
		// 值为1 表示之前出现过
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			// 之前未出现过，0->1,标记为出现过
			count[data[i]]++
		}
	}
	fmt.Println()
	//fmt.Println(count) // debug
}

func main() {
	//printRepeating(data)
	//printRepeating2(data)
	//printRepeating3(data)
	printRepeating4(data,4)
}
