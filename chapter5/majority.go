package main

// majority element ：出现次数超过数组大小一半的元素

import (
	"fmt"
	"sort"
)

//暴力：1.找到出现次数最多的元素；2.判断 次数是否过半
func getMajority(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0
	for i := 0; i < size; i++ {
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
	if maxCount > size/2 {
		return max, true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

// sort:1.排序；2.选中间索引元素为对比值；3.判断对比值出现次数是否过半
func getMajority2(data []int) (int, bool) {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data) // Sort(data,size)
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}

// 投票算法
func getMajority3(data []int) (int, bool) {
	majIndex := 0
	count := 1
	size := len(data)
	// 投票，选出票最多的
	for i := 1; i < size; i++ {
		// 相同 +1
		if data[majIndex] == data[i] {
			count++
		} else {
			// 不同 -1
			count--
		}

		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	fmt.Println(count, data[majIndex])
	// 计数 对比
	candidate := data[majIndex]
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}
func main() {
	data := []int{1, 2, 3, 1, 1, 2, 2, 3, 1, 1}
	fmt.Println(getMajority3(data))
}
