package main

import (
	"fmt"
	"sort"
)

/* 找到出现次数最多的元素*/
func getMax(data []int) (int, int) {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1
	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	return max, maxCount
}

// 先排序
func getMax2(data []int) int {
	size := len(data)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1
	sort.Ints(data) // Sort(data,size)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}

//Counting
func getMax3(data []int, dataRange int) int {
	max := data[0]
	maxCount := 1
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}
func main() {
	data := []int{1, 2, 3, 1, 1, 2, 2, 3, 1}
	fmt.Println(getMax(data))
}
