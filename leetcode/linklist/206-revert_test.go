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
