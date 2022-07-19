package _1_merge2linkedlist

import (
	"github.com/TomatoAres/algorithms-go/leetcode/linklist2"
	"testing"
)

type question21 struct {
	para21
	ans21
}

// para 是参数
// one 代表第一个参数
type para21 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans21 struct {
	one []int
}

func TestMerge(t *testing.T) {
	qs := []question21{

		{
			para21{[]int{}, []int{}},
			ans21{[]int{}},
		},
		//
		//{
		//	para21{[]int{1}, []int{1}},
		//	ans21{[]int{1, 1}},
		//},
		//
		//{
		//	para21{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		//	ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		//},
		//
		//{
		//	para21{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		//	ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		//},
		//
		//{
		//	para21{[]int{1}, []int{9, 9, 9, 9, 9}},
		//	ans21{[]int{1, 9, 9, 9, 9, 9}},
		//},
		//
		//{
		//	para21{[]int{9, 9, 9, 9, 9}, []int{1}},
		//	ans21{[]int{1, 9, 9, 9, 9, 9}},
		//},
		//
		//{
		//	para21{[]int{2, 3, 4}, []int{4, 5, 6}},
		//	ans21{[]int{2, 3, 4, 4, 5, 6}},
		//},
		//
		//{
		//	para21{[]int{1, 3, 8}, []int{1, 7}},
		//	ans21{[]int{1, 1, 3, 7, 8}},
		//},
		// 如需多个测试，可以复制上方元素。
	}

	t.Logf("------------------------Leetcode Problem 21------------------------\n")

	for _, ques := range qs {
		_, param := ques.ans21, ques.para21
		t.Logf("【input】:%v       【output】:%v\n", param,
			//linkList2.LL2Array(Merge2(linkList2.Array2LL(param.one), linkList2.Array2LL(param.another))))
			linkList2.LL2Array(MergeTwoLists(linkList2.Array2LL(param.one), linkList2.Array2LL(param.another))))
	}
	t.Logf("\n\n\n")
}
