package main

import "fmt"

func main() {
	root := &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}}
	subRoot := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}

	fmt.Println(isSubtree(root, subRoot))
}
