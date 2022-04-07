package linklist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	t.Logf("%#v\n", l)

	assert.Equal(t, New(), &LinkedList{})
	assert.NotNil(t, New())

}

var array1 = []int{1, 2, 3, 4, 5}
var link1 *LinkNode = &LinkNode{
	1, &LinkNode{
		2, &LinkNode{
			3, &LinkNode{
				4, &LinkNode{
					5, nil,
				}}}}}

//func Test_link2array(t *testing.T) {
//
//	assert.Equal(t, array1, link2array(link1))
//}
//
//func Test_Array2LinkedList(t *testing.T) {
//
//	// 指针值无法比较
//	//assert.Equal(t,link1,array2link(array1))
//
//	assert.Equal(t, array1, link2array(array2link(array1)))
//}

func Test_Travel(t *testing.T) {
	l := New()
	l.Travel()
}
