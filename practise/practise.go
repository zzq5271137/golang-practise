package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createTreeNode(v int) *treeNode {
	return &treeNode{v, nil, nil}
}

func main() {
	root := treeNode{2, nil, nil}
	root.left = &treeNode{3, &treeNode{2, nil, nil}, nil}
	root.right = createTreeNode(10)

	nodes := []treeNode{
		{2, nil, &root},
	}

	fmt.Println(nodes)
}
