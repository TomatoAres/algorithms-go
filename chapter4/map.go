package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[ ", key, " -> ", val, " ]")
	}
	// 即使这个键不存在，也不会报错
	fmt.Println("Apple price:", m["Apple"])
	delete(m, "Apple")
	// Apple 不存在， 默认输出 0 （int）
	fmt.Println("Apple price:", m["Apple"])
	v, ok := m["Apple"]
	fmt.Println("Apple price:", v, "Present:", ok)
	v2, ok2 := m["Banana"]
	fmt.Println("Banana price:", v2, "Present:", ok2)
	fmt.Println("keys=",len(m))

	m["boluo"] = 12
	fmt.Println("boluo price: ",m["boluo"])
}
