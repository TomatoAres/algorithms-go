package linkList2

import "fmt"

type LinkNode struct {
	Data int
	Next *LinkNode
}

// TODO 基本结构有问题
type LinkedList LinkNode

func New() *LinkedList {
	// new 申请空间, 这里返回的不是nil 了
	// Next 是nil，Next.value 必然报错
	return new(LinkedList)
}

func (l *LinkedList) IsEmpty() bool {
	return l.Next == nil
}

func (l *LinkedList) Clear() {
	l = New()
}

func (l *LinkedList) Travel() {
	if l.Next == nil {
		return
	}

	n := l.Next
	for n.Next != nil {
		fmt.Printf("%d\t", n.Data)
		n = n.Next
	}
	fmt.Println(n.Data)
}

func (l *LinkedList) Length() int {
	var length int

	//if l.Next == nil{
	//	return length
	//}

	p := l.Next
	for p != nil {
		length++
		p = p.Next
	}

	return length
}

func (l *LinkedList) Append(i int) {
	if l.Next == nil {
		l.Next = &LinkNode{
			Data: i,
		}
		return
	}

	p := l.Next
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &LinkNode{
		Data: i,
	}
	return
}
func Array2LL2(arr []int) *LinkedList {
	l := New()

	if len(arr) > 0 {
		node1 := &LinkNode{
			Data: arr[0],
		}

		p := node1
		for i := 1; i < len(arr); i++ {
			n := LinkNode{Data: arr[i]}
			p.Next = &n
			p = p.Next
		}
		// 添加头
		l.Next = p
	}

	return l
}
func Array2LL(arr []int) *LinkedList {
	l := New()

	for _, i := range arr {
		l.Append(i)
	}
	return l
}

func LL2Array(l *LinkedList) []int {
	var a []int
	if l.Next == nil {
		return a
	}

	p := l.Next
	for p != nil {
		a = append(a, p.Data)
		p = p.Next
	}
	return a
}
