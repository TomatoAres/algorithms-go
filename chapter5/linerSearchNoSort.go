package main

func SearchNoSort(src []int,e int)bool{
	for _,v := range src{
		if  v == e {
			return true
		}
	}
	return false
}

func SearchSort(src []int,e int)bool{
	for _,v := range src{
		if  v == e {
			return true
		}else if e < v{
			return false
		}

	}
	return false
}

func main()  {
	arr := make([]int,10)
	for i:=0;i<10;i++{
		arr[i] = i
	}

	println(SearchNoSort(arr,2))
	println(SearchSort(arr,2))
}