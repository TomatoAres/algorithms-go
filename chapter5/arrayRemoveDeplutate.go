package main

import (
	"fmt"
)

func RemoveDuplicatesAndEmpty(a []string) (ret []string){
	a_len := len(a)
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
			continue;
		}
		ret = append(ret, a[i])
	}
	return
}

//定义一个新切片（数组），存放原数组的第一个元素，
// 然后将新切片（数组）与原切片（数组）的元素一一对比，如果不同则存放在新切片（数组）中。
// 整体对比，需引入标志位 repeat
func RemoveDuplicates(a []string) (ret []string)  {
	ret = append(ret,a[0])
	a_len := len(a)
	for i:=1;i<a_len;i++{
		repeat := false
		for j:=0;j<len(ret);j++{
			if a[i] == a[j]{
				repeat = true
				break
			}
		}
		if !repeat{
			ret = append(ret,a[i])
		}
	}
	return
}
// i= 0 hello 有重复 repeat=true，不添加
// i= 1 " "   无重复 repeat=false 添加
// 出来还是乱序
// 参考：https://blog.csdn.net/benben_2015/article/details/80847966
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		// 引入标志 repeat
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}

		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func main() {
	a := []string{"hello", "", "world", "yes", "hello", "nihao", "shijie", "hello", "yes", "nihao", "good"}
	fmt.Println(RemoveDuplicates(a))
	////先排序了
	//sort.Strings(a)
	//fmt.Println(a)
	//fmt.Println(RemoveDuplicatesAndEmpty(a))
}
