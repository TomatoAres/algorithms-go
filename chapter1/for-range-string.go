package main

import "fmt"

func main() {
	s := "hello"
	for i,c := range s{
		// 默认 ASCII 值
		fmt.Printf("%v,%v,%c\n",i,c,c)
	} 
}