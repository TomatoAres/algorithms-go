package main

import "fmt"

func main() {
	s := "hello World"
	r := []rune(s)
	fmt.Printf("before:\ts=%s,pointer=%p\n",s,&s)

	// 单个字符需要使用 单引号 ''
	r[0]='H'
	s = "Hello world"

	// 不可变的意思是说 你不能像c语言那样去修改字符串的值， string可以理解成字符slice一样，只是不能修改
	// s[0] = 'H'  // cannot assign to s[0]

	// 地址值没有变的原因：声明的就是一个指针，指向的地方实际上已经变了
	fmt.Printf("after:\ts=%s,pointer=%p\n",s,&s)

	fmt.Println(string(r))
}