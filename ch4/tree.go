package main

import "fmt"

type point struct {
	i, j int
}

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() { // print作为node的接收者
	fmt.Print(node.value, "")
}

// 指针接收者
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored.")
		return
	}
	node.value = value
}

// 创建工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value} // 局部变量也可以返回给别人用
}

func main() {
	var root treeNode
	root = treeNode{value: 3} // 前面没写var 可写成 root :=treeNode{}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // go中不管是指针还是实例都是点.下去
	root.left.right = createNode(2)
	//fmt.Println(root)               // {0 <nil> <nil>}
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	root.print()
	fmt.Println(nodes)

	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()
}
