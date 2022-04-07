package linkList2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	assert.NotNil(t, l)
	assert.Equal(t, 0, l.Data)
	assert.Nil(t, l.Next)
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, New().IsEmpty())
}

func TestLength(t *testing.T) {
	assert.Equal(t, 0, New().Length())
}

func TestArray2LL(t *testing.T) {
	//arr := []int{}
	assert.Equal(t, New(), Array2LL([]int{}))
	assert.Equal(t,
		&LinkedList{Next: &LinkNode{Data: 0}},
		Array2LL([]int{0}))
}

func TestLL2Array(t *testing.T) {
	assert.Nil(t, LL2Array(New()), "empty slice is nil")

	// alice 是 指针，指针不能比较
	//assert.Equal(t, LL2Array(&LinkedList{Next: &LinkNode{Data: 0}}),[]int{0})

	result := LL2Array(&LinkedList{Next: &LinkNode{Data: 0}})
	expected := []int{0}
	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected[0], result[0])
}

// TODO: slice 坑
func TestSlice(t *testing.T) {
	var a []int     // a == nil
	var b = []int{} // b != nil

	t.Logf("%#v\n", a)
	t.Logf("%#v\n", b)

	//assert.Equal(t, a,nil)
	assert.Nil(t, a) // pass
	//assert.Equal(t, nil,a) // 应该用 assert.Nil
	//assert.Nil(t, b)
}

func TestTravel(t *testing.T) {
	s := []int{1, 2, 3}
	l := Array2LL2(s)
	t.Log(l.Length())
	l.Travel()

	t.Logf("%#v", LL2Array(l))
}
