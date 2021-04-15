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

type Node struct {
	data       int
	left       *Node
	right      *Node
	dfsDataArr []int
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{data: data}
		return
	}
	t.root.Insert(data)
}

func (n *Node) Insert(data int) {
	if n.data == 0 {
		n.data = data
		return
	}
	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.Insert(data)
		}
		return
	}

	if n.right == nil {
		n.right = &Node{data: data}
	} else {
		n.right.Insert(data)
	}
}

func (n *Node) Contains(data int) bool {
	if n.data == data {
		return true
	}

	shouldLookLeft := data < n.data
	if shouldLookLeft {
		if n.left == nil {
			return false
		}
		return n.left.Contains(data)
	}

	// look right
	if n.right == nil {
		return false
	}
	return n.right.Contains(data)
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
		return n.data
	}
	dfsDataArr := n.GetBFSDataArray()
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
	if len(dfsDataArr)-1 < dataArrIndex {
		return 0
	}
	return dfsDataArr[dataArrIndex]
}

func (n *Node) GetBFSDataArray() []int {
	if len(n.dfsDataArr) > 0 {
		// cache
		return n.dfsDataArr
	}

	res := make([]int, 0)
	res = append(res, n.data)
	queue := make([]*Node, 0)
	queue = append(queue, n)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// add children to queue
		if node.left != nil {
			queue = append(queue, node.left)
			res = append(res, node.left.data)
		} else {
			res = append(res, 0)
		}
		if node.right != nil {
			queue = append(queue, node.right)
			res = append(res, node.right.data)
		} else {
			res = append(res, 0)
		}
	}

	n.dfsDataArr = res // add cache
	return res
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
