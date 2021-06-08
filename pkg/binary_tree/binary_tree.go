package binary_tree

import (
	"container/list"
	"fmt"
)

type BinaryTree struct {
	Nodes []Node
	Head  *Node
}

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type QueueItem struct {
	Node   *Node
	Middle int
	Start  int
	End    int
}

func GenerateBinaryTreeFromSortedArray(a []int) BinaryTree {

	n := len(a)
	visited := make([]int, n)

	bt := BinaryTree{
		Nodes: []Node{},
	}

	for i := 0; i < n; i++ {
		bt.Nodes = append(bt.Nodes, Node{
			Value: a[i],
		})
	}

	queue := list.New()
	index := n / 2

	queue.PushBack(QueueItem{
		Node:   &bt.Nodes[index],
		Middle: index,
		Start:  0,
		End:    n,
	})

	bt.Head = &bt.Nodes[index]

	for queue.Len() > 0 {
		e := queue.Front()
		item := e.Value.(QueueItem)
		curr := item.Node
		index := item.Middle
		start := item.Start
		end := item.End

		visited[index] = 1

		indexLeft := (index-start)/2 + start
		if indexLeft == index || visited[indexLeft] == 1 {
			curr.Left = nil
		} else {
			curr.Left = &bt.Nodes[indexLeft]
			curr.Left.Parent = curr
			queue.PushBack(QueueItem{
				Node:   curr.Left,
				Middle: indexLeft,
				Start:  start,
				End:    index,
			})
		}

		indexRight := (end-index)/2 + index
		if indexRight == index || visited[indexRight] == 1 {
			curr.Right = nil
		} else {
			curr.Right = &bt.Nodes[indexRight]
			curr.Right.Parent = curr
			queue.PushBack(QueueItem{
				Node:   curr.Right,
				Middle: indexRight,
				Start:  index,
				End:    end,
			})
		}

		queue.Remove(e)

	}

	return bt
}

func GetHeight(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := GetHeight(node.Left)
	rightHeight := GetHeight(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (bt *BinaryTree) Traverse() {
	traverse(bt.Head)
}

func (bt *BinaryTree) Search(value int) *Node {
	return search(bt.Head, value)
}

func (bt *BinaryTree) Insert(value int) {
	insert(bt.Head, value, nil, bt)
}

func search(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

func traverse(node *Node) {
	if node != nil {
		traverse(node.Left)
		fmt.Print(node.Value)
		if node.Left == nil && node.Right == nil {
			fmt.Println(" is a leaf")
		} else {
			fmt.Println()
		}
		traverse(node.Right)
	}
}

func insert(node *Node, value int, parent *Node, bt *BinaryTree) {

	if node != nil {
		if value < node.Value {
			insert(node.Left, value, node, bt)
		} else {
			insert(node.Right, value, node, bt)
		}
	} else {
		newNode := Node{
			Value:  value,
			Parent: parent,
			Left:   nil,
			Right:  nil,
		}
		bt.Nodes = append(bt.Nodes, newNode)
		if parent.Value > value {
			parent.Left = &bt.Nodes[len(bt.Nodes)-1]
		} else {
			parent.Right = &bt.Nodes[len(bt.Nodes)-1]
		}
	}
}
