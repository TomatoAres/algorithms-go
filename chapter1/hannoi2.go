package main

import "fmt"

func main() {
	// 数字是初始柱子上的盘子数
	TowersOfHanoi(2)
}
/*
要将N个磁盘从源移动到目标，
1.我们首先将N-1个磁盘从源移动到临时磁盘，
2. 然后将最低的N个磁盘从源移动到目标。
3.然后将N-1个磁盘从临时磁盘移动到目标磁盘
from	源
to		目的
temp	临时
*/
func TOHUtil(num int, from string, to string, temp string) {
	// 没有盘子，停止
	if num < 1 {
		return
	}
	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	TOHUtil(num-1, temp, to, from)
}

/* TOH=tower of hanoi */
func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	TOHUtil(num, "A", "C", "B")
}
