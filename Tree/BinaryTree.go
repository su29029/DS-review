package main

import (
	"fmt"
	"reflect"
)

//BinaryTree the binary tree structure
type BinaryTree struct {
	val   interface{}
	left  *BinaryTree
	right *BinaryTree
}

// CreateAnyTypeSlice converse interface{} to slice
func CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := isSlice(slice)
	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

func isSlice(slice interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(slice)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

// CreateByPreOrderSequence create a binary tree by a preorder sequence(which contains node and nil, nil means an empty subtree)
func CreateByPreOrderSequence(arr interface{}) *BinaryTree {
	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		fmt.Errorf("not a slice")
		return nil
	}
	if len(slice) == 0 {
		return new(BinaryTree)
	}

	var create func([]interface{}, *int) *BinaryTree
	create = func(slice []interface{}, flag *int) *BinaryTree {
		if *flag < len(slice)-1 {
			*flag = *flag + 1
		}
		if slice[*flag] == nil {
			return nil
		}
		node := new(BinaryTree)
		node.val = slice[*flag]
		node.left = create(slice, flag)
		node.right = create(slice, flag)

		return node
	}

	flag := -1
	binaryTree := create(slice, &flag)

	// fmt.Println(binaryTree.TraversePreOrder())
	return binaryTree
}

// CreateByTwoSequences create a binary tree by two sequences(preorder sequence/inorder sequence which flag is 1, and inorder sequence/postorder sequence which flag is 2)
func (binaryTree *BinaryTree) CreateByTwoSequences(seq1 interface{}, seq2 interface{}, flag int) error {
	return nil
}

// TraversePreOrder traverse in preorder sequence
func (binaryTree *BinaryTree) TraversePreOrder() interface{} {

	var traverse func(root *BinaryTree)
	var res []interface{}
	traverse = func(root *BinaryTree) {
		if root == nil {
			return
		}
		res = append(res, root.val)
		traverse(root.left)
		traverse(root.right)
	}

	traverse(binaryTree)
	return res
}

// TraverseInOrder traverse in inorder sequence
func (binaryTree *BinaryTree) TraverseInOrder() interface{} {
	var traverse func(root *BinaryTree)
	var res []interface{}
	traverse = func(root *BinaryTree) {
		if root == nil {
			return
		}

		traverse(root.left)
		res = append(res, root.val)
		traverse(root.right)
	}

	traverse(binaryTree)
	return res
}

// TraversePostOrder traverse in postorder sequence
func (binaryTree *BinaryTree) TraversePostOrder() interface{} {

	var traverse func(root *BinaryTree)
	var res []interface{}
	traverse = func(root *BinaryTree) {
		if root == nil {
			return
		}
		traverse(root.left)
		traverse(root.right)
		res = append(res, root.val)
	}

	traverse(binaryTree)
	return res
}
