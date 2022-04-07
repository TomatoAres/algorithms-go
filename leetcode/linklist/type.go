package linklist

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}

// 太麻烦
type LinkedList struct {
	head *LinkNode
}

// New return &LinkedList
// 有头结点的
func New() *LinkedList {
	// 这里返回的不是nil 了
	// head 是nil，head.value 必然报错
	return &LinkedList{}
}

func (l *LinkedList) Travel() {
	if l.head != nil {
		fmt.Printf("%d\t", l.head.data)
	}
	fmt.Println()
	//fmt.Println(l.head.data)
	fmt.Println(l.head.next)
}

func (l *LinkedList) Insert(i int) {
	n := &LinkNode{data: i}

	p := l.head
	l.head.next = n
	n.next = p

}

func (l *LinkedList) Append(i int) {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = &LinkNode{data: i}
}

func Array2LinkedList(array []int) *LinkedList {
	l := New()

	for _, i := range array {
		node := LinkNode{
			data: i,
		}
		//l.Insert(i)
		l.head.next = &node
	}

	return l
}
