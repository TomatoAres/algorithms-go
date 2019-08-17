package main

import "fmt"

// 递归出口和循环停止条件一致
func BinarySearch(array []int,e int)bool  {
	start := 0
	end := len(array)-1
	
	mid := (start+end)/2
	if start >= end{
		return false
	}
	
	if e == array[mid]{
		return true
	} else if e < array[mid]{
		return BinarySearch(array[:mid-1],e)
	}else {
		return BinarySearch(array[mid+1:],e)
	}
}

func BinarySearchFor(array []int,e int)bool  {
	start := 0
	end := len(array)-1
	var mid int
	for start <=end {
		mid = (start + end)/2
		if array[mid]==e{
			return true
		}else if e < array[mid]{
			end = mid -1
		}else {
			start = mid +1
		}
	}
	return false
}

func main() {
	data := []int{1,2,3,4,5,6}
	value := 0
	fmt.Println(BinarySearchFor(data,value))
	fmt.Println(BinarySearch(data,value))
}
