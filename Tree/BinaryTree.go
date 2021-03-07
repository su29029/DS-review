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

	return binaryTree
}

// CreateByTwoSequences create a binary tree by two sequences(preorder sequence/inorder sequence which flag is 1, and inorder sequence/postorder sequence which flag is 2)
func CreateByTwoSequences(seq1 interface{}, seq2 interface{}, flag int) *BinaryTree {
	slice1, ok1 := CreateAnyTypeSlice(seq1)
	slice2, ok2 := CreateAnyTypeSlice(seq2)
	if !ok1 || !ok2 {
		fmt.Errorf("not slices")
		return nil
	}

	if len(slice1) != len(slice2) {
		fmt.Errorf("invalid parameters")
		return nil
	}

	var createByPreAndIn func([]interface{}, int, int, int, map[interface{}]int) *BinaryTree
	var createByInAndPost func([]interface{}, int, int, int, map[interface{}]int) *BinaryTree

	var binaryTree = new(BinaryTree)
	var inPos map[interface{}]int
	inPos = make(map[interface{}]int)

	createByPreAndIn = func(preOrder []interface{}, preStart int, preEnd int, inStart int, inPos map[interface{}]int) *BinaryTree {
		if preStart > preEnd {
			return nil
		}
		node := new(BinaryTree)
		node.val = preOrder[preStart]
		rootIdx := inPos[preOrder[preStart]]
		leftLen := rootIdx - inStart

		node.left = createByPreAndIn(preOrder, preStart+1, preStart+leftLen, inStart, inPos)
		node.right = createByPreAndIn(preOrder, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
		return node
	}

	createByInAndPost = func(postOrder []interface{}, postStart int, postEnd int, inStart int, inPost map[interface{}]int) *BinaryTree {
		if postStart > postEnd {
			return nil
		}

		node := new(BinaryTree)
		node.val = postOrder[postEnd]
		rootIdx := inPos[postOrder[postEnd]]
		leftLen := rootIdx - inStart

		node.left = createByInAndPost(postOrder, postStart, postStart+leftLen-1, inStart, inPos)
		node.right = createByInAndPost(postOrder, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
		return node
	}

	// save the inOrder's node, instead of traversing the inOrder sequence in the recursion
	for i := 0; i < len(slice1); i++ {
		inPos[slice1[i]] = i
	}

	if flag == 1 {
		binaryTree = createByPreAndIn(slice1, 0, len(slice1)-1, 0, inPos)
	}
	if flag == 2 {
		binaryTree = createByInAndPost(slice2, 0, len(slice2)-1, 0, inPos)
	}

	return binaryTree
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

// IterTraversePreOrder Iteratively traverse binaryTree in preorder sequence
func (binaryTree *BinaryTree) IterTraversePreOrder() interface{} {
	var res []interface{}
	var stack []*BinaryTree

	root := binaryTree
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.val)
			stack = append(stack, root)
			root = root.left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.right
		}
	}
	return res
}

// IterTraverseInOrder Iteratively traverse binaryTree in inorder sequence
func (binaryTree *BinaryTree) IterTraverseInOrder() interface{} {
	var res []interface{}
	var stack []*BinaryTree

	root := binaryTree
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.val)
			root = root.right
		}
	}
	return res
}

// IterTraversePostOrder Iteratively traverse binaryTree in postorder sequence
func (binaryTree *BinaryTree) IterTraversePostOrder() interface{} {
	var res []interface{}
	var stack []*BinaryTree

	var cur *BinaryTree
	var pre *BinaryTree = nil

	root := binaryTree
	stack = append(stack, root)
	for len(stack) != 0 {
		cur = stack[len(stack)-1]

		if (cur.left == nil && cur.right == nil) || (pre != nil && (pre == cur.left || pre == cur.right)) {
			res = append(res, cur.val)
			stack = stack[:len(stack)-1]
			pre = cur
		} else {
			if cur.right != nil {
				stack = append(stack, cur.right)
			}
			if cur.left != nil {
				stack = append(stack, cur.left)
			}
		}
	}
	return res
}
