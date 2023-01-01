package dsalgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	stack := Stack{}

	Convey("Test Stack Pop", t, func() {
		Convey("should add the first node", func() {
			newNode := Node{}
			newNode.SetValue(10)
			stack.Push(&newNode)
			resultNode := stack.Peek()
			assert.Equal(t, 10, resultNode.GetValue())
		})

		Convey("should add the 2nd node", func() {
			newNode2 := Node{}
			newNode2.SetValue(20)
			stack.Push(&newNode2)
			resultNode := stack.Peek()
			assert.Equal(t, 20, resultNode.GetValue())
		})
	})
}

func TestStackPop(t *testing.T) {
	stack := Stack{}

	Convey("Test Stack Pop", t, func() {
		Convey("should return nil if stack is empty", func() {
			assert.Nil(t, stack.Pop())
		})

		Convey("should return the last node if stack only has 1 node", func() {
			newNode := Node{}
			newNode.SetValue(10)
			stack.Push(&newNode)
			assert.Equal(t, 1, stack.Len())
			poppedNode := stack.Pop()
			assert.Equal(t, 10, poppedNode.GetValue())
			assert.Equal(t, 0, stack.Len())
			lastNode := stack.Peek()
			assert.Nil(t, lastNode)
		})

		Convey("should return the last node if stack has more than 1 nodes", func() {
			newNode1 := Node{}
			newNode1.SetValue(1)
			stack.Push(&newNode1)
			newNode2 := Node{}
			newNode2.SetValue(2)
			stack.Push(&newNode2)
			assert.Equal(t, 2, stack.Len())
			poppedNode := stack.Pop()
			assert.Equal(t, 2, poppedNode.GetValue())
		})
	})
}
