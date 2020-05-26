package linklist

type LinkNode struct {
	data int
	next *LinkNode
}

func array2link(array []int) *LinkNode {
	if len(array) == 0 {
		return nil
	}

	// 构造头节点--链表
	var link *LinkNode
	link = &LinkNode{
		data: array[0],
		next: nil,
	}

	// 不得不引入 tmp
	tmp := link
	//var nextNode *LinkNode
	for i := 1; i < len(array); i++ {
		nextNode := &LinkNode{
			data: array[i],
			next: nil,
		}
		tmp.next = nextNode
		tmp = tmp.next
	}

	return link
}

func link2array(link *LinkNode) []int {
	// 垃圾
	//if link == nil{
	//	return nil
	//}

	var result []int
	for link != nil {
		result = append(result, link.data)
		link = link.next
	}
	return result
}
