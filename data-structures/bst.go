package dsalgo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	LEFT  = 0
	RIGHT = 1
)

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}
	t.root.Insert(value)
}

func (n *Node) Insert(value int) {
	if n.value == 0 {
		n.value = value
		return
	}
	if value < n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.Insert(value)
		}
		return
	}

	if n.right == nil {
		n.right = &Node{value: value}
	} else {
		n.right.Insert(value)
	}
}

func (n *Node) Contains(value int) bool {
	if n.value == value {
		return true
	}

	shouldLookLeft := value < n.value
	if shouldLookLeft {
		if n.left == nil {
			return false
		}
		return n.left.Contains(value)
	}

	// look right
	if n.right == nil {
		return false
	}
	return n.right.Contains(value)
}

/*
print a node like this
		  5
	 4	      9
   1   2    6   12
*/
func (n *Node) Print() {
	height := n.GetHeight()
	fmt.Println("height: ", height)
	numOfLeaves := int(math.Pow(2, float64(height-1)))
	w := numOfLeaves * (2 + 4 + 2) // left margin + num + right margin
	fmt.Println("width: ", w)
	fmt.Println("numOfLeaves ", numOfLeaves)

	for i := 0; i < height; i++ {
		line := ""
		iNumOfLeaves := int(math.Pow(2, float64(i)))
		iNumOfSpaces := iNumOfLeaves + 1
		spaceWidth := int((w - 4*iNumOfLeaves) / iNumOfSpaces)
		// fmt.Printf("h %v, spaceWidth: %v", height, spaceWidth)
		if i == 0 {
			spaceWidth = int(w/2) + 1
		}

		for j := 0; j < iNumOfLeaves; j++ {
			numberStr := ""
			number := n.GetNodeDataByHeightAndIndex(i+1, j)
			if number != 0 {
				numberStr = strconv.Itoa(number)
			} else {
				numberStr = "n"
			}
			if len(numberStr) < 4 {
				numberStr = strings.Repeat(" ", 4-len(numberStr)) + numberStr
			}
			line = line + strings.Repeat(" ", spaceWidth) + numberStr + strings.Repeat(" ", spaceWidth)
		}
		fmt.Printf("%s\n", line)
	}
}

func (n *Node) GetNodeDataByHeightAndIndex(h int, index int) int {
	if h == 1 {
		return n.value
	}
	bfsDataArr := n.GetBFSDataArrayWithZeroFills()
	numOfNodesByPrevLevel := 1
	currNumOfNodes := 2
	for i := 1; i < h; i++ {
		if i == 1 {
			continue
		}
		numOfNodesByPrevLevel = numOfNodesByPrevLevel + currNumOfNodes
		currNumOfNodes = currNumOfNodes * 2
	}

	dataArrIndex := (numOfNodesByPrevLevel - 1) + index + 1
	if len(bfsDataArr)-1 < dataArrIndex {
		return 0
	}
	return bfsDataArr[dataArrIndex]
}

func (n *Node) GetBFSDataArrayWithZeroFills() []int {
	if len(n.bfsDataArr) > 0 {
		// cache
		return n.bfsDataArr
	}

	res := make([]int, 0)
	res = append(res, n.value)
	queue := QueueForNode{}
	queue.Enqueue(n)

	for queue.Len() > 0 {
		node, _ := queue.Dequeue()
		// add children to queue
		if node.left != nil {
			queue.Enqueue(node.left)
			res = append(res, node.left.value)
		} else {
			res = append(res, 0)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
			res = append(res, node.right.value)
		} else {
			res = append(res, 0)
		}
	}

	n.bfsDataArr = res // add cache
	return res
}

func (n *Node) BreathFirstSearch() []int {
	if len(n.bfsDataArr) > 0 {
		// cache
		return n.bfsDataArr
	}

	res := make([]int, 0)
	res = append(res, n.value)
	queue := QueueForNode{}
	queue.Enqueue(n)

	for queue.Len() > 0 {
		node, _ := queue.Dequeue()
		// add children to queue
		if node.left != nil {
			queue.Enqueue(node.left)
			res = append(res, node.left.value)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
			res = append(res, node.right.value)
		}
	}

	n.bfsDataArr = res // add cache
	return res
}

func (n *Node) DepthFirstSearch() []int {
	// cache
	if len(n.dfsDataArr) > 0 {
		return n.dfsDataArr
	}

	list := make([]int, 0)
	list = append(list, n.GetValue())
	currentLeftNode := n.GetLeft()
	list = append(list, currentLeftNode.GetValue())
	for currentLeftNode != nil {
		if currentLeftNode.GetLeft() != nil {
			list = append(list, currentLeftNode.GetLeft().GetValue())
		}
		if currentLeftNode.GetRight() != nil {
			list = append(list, currentLeftNode.GetRight().GetValue())
		}
		currentLeftNode = currentLeftNode.GetLeft()
	}

	currentRightNode := n.GetRight()
	list = append(list, currentRightNode.GetValue())
	for currentRightNode != nil {
		if currentRightNode.GetLeft() != nil {
			list = append(list, currentRightNode.GetLeft().GetValue())
		}
		if currentRightNode.GetRight() != nil {
			list = append(list, currentRightNode.GetRight().GetValue())
		}
		currentRightNode = currentRightNode.GetRight()
	}

	n.dfsDataArr = list
	return list
}

func DepthFirstSearchRecursive(node *Node, list []int) []int {
	if node == nil {
		return list
	}
	if node.GetLeft() != nil {
		list = append(list, DepthFirstSearchRecursive(node.GetLeft(), list)...)
	}
	if node.GetRight() != nil {
		list = append(list, DepthFirstSearchRecursive(node.GetRight(), list)...)
	}
	list = append(list, node.GetValue())
	return list
}

func (n *Node) GetHeight() int {
	if n == nil {
		return 0
	}

	leftHeight := n.left.GetHeight()
	rightHeight := n.right.GetHeight()

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func BuildNodeFromSlice(numArr []int) *Node {
	var node = Node{}
	for _, i := range numArr {
		node.Insert(i)
	}
	return &node
}
