package main

import "fmt"
var count int
func PrintSlice(array []int) {
	count += 1
	fmt.Printf("%d data:",count)
	fmt.Println(array)
}

// TODO
//核心：递归出口 i==length；过程 交换-递归-交换；
func Permutation(data []int, i int, length int) {
	if length == i {
		PrintSlice(data)
		return
	}
	for j := i; j < length; j++ {
		//交换
		swap(data, i, j)
		//递归
		Permutation(data, i+1, length)
		//交换
		swap(data, i, j)
	}
}

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func main() {
	// 构建初始数据
	var data [5]int
	for i := 0; i < 5; i++ {
		data[i] = i
	}
	// 全排列
	Permutation(data[:], 0, 5)
}
