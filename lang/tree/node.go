package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// Print 给结构体增加print方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue  给结构体增加setValue方法
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

// CreateNode 创建Node 工厂
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
