package main

//Find the largest sum contiguous subarray.

import "fmt"
/*
思路：
1.基本思路就是遍历一遍，用两个变量，一个记录最大的和，一个记录当前的和。
时空复杂度貌似还不错......（时间复杂度 O(n)O(n)，空间复杂度 O(l)O(l)）

参考链接：https://leetcode-cn.com/problems/two-sum/solution/bao-li-qiu-jie-by-pandawakaka/

*/
// TODO 暂时不能理解，先跳过
func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0
	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func main() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Println("Max sub array sum :", MaxSubArraySum(data))
}
