package main

import (
	"fmt"
	"tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.TreeNode
	root = tree.Node{Value: 3} // 如果前面没写var root, 可写成 root :=treeNode{}
	root.left = &tree.Node{}   // left是一个指针,取地址
	root.right = &tree.Node{5, nil, nil}
	root.right.left = new(tree.Node) // go中不管是指针还是实例都是点.下去
	root.left.right = tree.CreateNode(2)
	root.right.left.setValue(4)
	root.right.left.print()

	root.traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

}
