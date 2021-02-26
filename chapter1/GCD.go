package main

import "fmt"

//Greatest common divisor (GCD)
// 默认 m>n
func GCD(m, n int) int {
	if m < n {
		return GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	// 第一个参数 取m，n都行，最好取小一点，更快
	return GCD(n, m%n)
}

func main() {
	fmt.Println(GCD(12, 3))
}
