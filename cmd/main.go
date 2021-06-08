package main

import (
	"fmt"
	bt "github.com/alessandrobessi/algorithms/pkg/binary_tree"
)

func main() {
	fmt.Println("generate a binary tree from a sorted array...")
	binaryTree := bt.GenerateBinaryTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 8, 9, 10})
	binaryTree.Traverse()

	fmt.Println("\n\nget the height of the binary tree")
	fmt.Println(bt.GetHeight(binaryTree.Head))

	fmt.Println("\n\ninsert 7")
	binaryTree.Insert(7)
	binaryTree.Traverse()

	fmt.Println("\n\nsearch a node containing the value 6...")
	node := binaryTree.Search(6)
	fmt.Println(node)

	fmt.Println("\n\nsearch a node containing the value 100...")
	node = binaryTree.Search(100)
	fmt.Println(node)

	fmt.Println("\n\nget the height of the binary tree")
	fmt.Println(bt.GetHeight(binaryTree.Head))
}
