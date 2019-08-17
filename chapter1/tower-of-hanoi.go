package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("输入要移动的盘子数：")
	reader := bufio.NewReader(os.Stdin)
lool:
	data, _, _ := reader.ReadLine()
	n, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Println(err)
		goto lool
	}
	hanoi(n, 'A', 'B', 'C')
}

func hanoi(n int, a, b, c byte) {
	if n > 1 {
		hanoi(n-1, a, c, b)
		fmt.Printf("%c-->%c\n", a, c)
		hanoi(n-1, b, a, c)
	} else {
		fmt.Printf("%c-->%c\n", a, c)
	}
}
