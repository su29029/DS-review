package main

import "fmt"

func main() {
	fmt.Println("Tree")
	// bTree := new(BinaryTree)
	bTree := CreateByPreOrderSequence([]interface{}{1, 2, 4, nil, 7, nil, nil, nil, 3, 5, nil, nil, 6, nil})
	fmt.Println(bTree.TraversePreOrder())
	fmt.Println(bTree.TraverseInOrder())
	fmt.Println(bTree.TraversePostOrder())
}
