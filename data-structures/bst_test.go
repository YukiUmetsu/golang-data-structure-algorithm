package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeInsert(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 17})
	assert.Equal(t, node.left.value, 5)
	assert.Equal(t, node.right.value, 15)
	assert.Equal(t, node.right.right.value, 17)
}

func TestNodeContains(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 20, 0, -5, 3})
	assert.Equal(t, true, node.Contains(3))
	assert.Equal(t, false, node.Contains(-100))
}

func TestNodePrint(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 20, -5, 3})
	node.Print()
	// fmt.Println("")
	// node2 := BuildNodeFromSlice([]int{10, 5, 15, 17})
	// node2.Print()
}

func TestGetBFSDataArray(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 20, -5, 3})
	list := node.BreathFirstSearch()
	assert.Equal(t, len(list), 6)
	assert.Equal(t, []int{10, 5, 15, -5, 20, 3}, list)
}

func TestDepthFirstSearch(t *testing.T) {
	node := BuildNodeFromSlice([]int{10, 5, 15, 20, -5, 3})
	list := node.DepthFirstSearch()
	assert.Equal(t, len(list), 6)
	assert.Equal(t, []int{10, 5, -5, 3, 15, 20}, list)

	node2 := BuildNodeFromSlice([]int{9, 4, 6, 20, 170, 15, 1})
	list2 := node2.DepthFirstSearch()
	assert.Equal(t, len(list2), 7)
	assert.Equal(t, []int{9, 4, 1, 6, 20, 15, 170}, list2)
}
