package main

import (
	"fmt"
	"reflect"
)

// LinkedNode the base data structure
type LinkedNode struct {
	val  interface{}
	next *LinkedNode
}

// LinkedList the alias of LinkedNode
type LinkedList = LinkedNode

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

// check if a slice
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

// Create create a LinkedList
func (linkedList *LinkedNode) Create(arr interface{}) (err error) {
	// converse interface{} to slice
	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		return fmt.Errorf("not a slice")
	}

	if len(slice) == 0 {
		return fmt.Errorf("slice's length is 0")
	}

	linkedList.val = slice[0]

	pNode := new(LinkedNode)
	linkedList.next = pNode

	for i, value := range slice {
		// jump off the first node
		if i == 0 {
			continue
		}
		pNode.val = value

		// not the last node
		if i < len(slice)-1 {
			pNode.next = new(LinkedNode)
			pNode = pNode.next
		}
	}

	return
}

// Print print all nodes' value in the linkedlist
func (linkedList *LinkedNode) Print() {
	for node := linkedList; node != nil; node = node.next {
		fmt.Print(node.val, " ")
	}
	fmt.Print("\n")
}

// Add add a node at the end of the linkedlist
func (linkedList *LinkedNode) Add(val interface{}) (err error) {
	node := linkedList
	for node.next != nil {
		node = node.next
	}
	node.next = new(LinkedNode)
	node.next.val = val
	node.next.next = nil

	return
}

// Modify modify a node in the linkedlist
func (linkedList *LinkedNode) Modify(oldVal interface{}, newVal interface{}, pos int) (err error) {
	node := linkedList
	var i int = 0 // the i-th element that satisfies the condition
	for node != nil {
		if node.val == oldVal {
			i++
			if i == pos {
				node.val = newVal
				return
			}
		}
		node = node.next
	}
	return fmt.Errorf("cannot find the %v-th value %v in the linkedlist", pos, oldVal)
}

// Search search if a value in the linkedlist, if so, return true, if not, return false
func (linkedList *LinkedNode) Search(value interface{}, pos int) (ok bool, err error) {
	node := linkedList
	var i int = 0 // the count
	for node != nil {
		if node.val == value {
			i++
			if i == pos {
				ok = true
				break
			}
		}
		node = node.next
	}
	return false, fmt.Errorf("not found %v-th value %v", pos, value)
}

// Delete delete a node of this linkedlist
func (linkedList *LinkedNode) Delete(value interface{}, pos int) (ok bool, err error) {
	node := linkedList
	var i int = 0 // the count
	for node != nil {
		if node.next.val == value {
			i++
			if i == pos {
				// delete the node
				node.next = node.next.next
				ok = true
			}
		}
		node = node.next
	}
	return false, fmt.Errorf("not found %v-th value %v", pos, value)
}
