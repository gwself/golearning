package main

import (
	"fmt"
	"golearning/lang/tree"
)

func main() {
	testNode()
}

func testNode() {
	root := tree.Node{Value: 7}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 3, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(5)
	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()
	root.Traverse()
}
