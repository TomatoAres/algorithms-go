package linklist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var revertQ = []int{1, 2, 3, 4, 5}
var revertA = []int{5, 4, 3, 2, 1}

var revertNil []int

func Test_revert(t *testing.T) {

	assert.Equal(t, revertA, link2array(revert(array2link(revertQ))))

	assert.Equal(t, revertNil, link2array(revert(array2link(revertNil))))
}
func TestNew(t *testing.T) {
	assert.Equal(t, (*LinkNode)(nil), New())
}

func Benchmark_revert(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		revert(array2link(revertQ))
	}
}

func Benchmark_array2link(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		array2link(revertQ)
	}
}

func Test_revertRecursion(t *testing.T) {
	assert.Equal(t, revertNil, link2array(revertRecursion(array2link(revertNil))))

	assert.Equal(t, revertA, link2array(revertRecursion(array2link(revertQ))))
}

func Benchmark_revertRecursion(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		revertRecursion(array2link(revertQ))

	}
}
