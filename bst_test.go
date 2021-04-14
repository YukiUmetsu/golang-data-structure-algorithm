package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeInsert(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 17})
	assert.Equal(t, node.left.data, 5)
	assert.Equal(t, node.right.data, 15)
	assert.Equal(t, node.right.right.data, 17)
}

func TestNodeContains(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 20, 0, -5, 3})
	assert.Equal(t, true, node.Contains(3))
	assert.Equal(t, false, node.Contains(-100))
}

func TestNodePrint(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 20, 0, -5, 3})
	node.Print()
	// fmt.Println("")
	// node2 := BuildNodeFromSlice([]int{10, 5, 15, 17})
	// node2.Print()
}
