package main

import "fmt"

func main() {
	fmt.Println("Tree")
	binaryTree := CreateByPreOrderSequence([]interface{}{1, 2, 4, nil, 7, nil, nil, nil, 3, 5, nil, nil, 6, nil})
	fmt.Println(binaryTree.TraversePreOrder())
	fmt.Println(binaryTree.TraverseInOrder())
	fmt.Println(binaryTree.TraversePostOrder())
	binaryTree2 := CreateByTwoSequences([]interface{}{1, 2, 4, 5, 3, 6, 7, 8}, []interface{}{4, 2, 5, 1, 7, 6, 8, 3}, 1)
	fmt.Println(binaryTree2.TraversePreOrder())
	binaryTree3 := CreateByTwoSequences([]interface{}{4, 2, 5, 1, 7, 6, 8, 3}, []interface{}{4, 5, 2, 7, 8, 6, 3, 1}, 2)
	fmt.Println(binaryTree3.TraversePreOrder())
	fmt.Println(binaryTree3.IterTraversePreOrder())
	fmt.Println(binaryTree3.IterTraverseInOrder())
	fmt.Println(binaryTree3.IterTraversePostOrder())
}
