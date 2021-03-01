package main

import (
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

// 判断是否为slcie数据
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

// Create create a LinkedList
func (linkedList *LinkedNode) Create(arr interface{}) {
	// converse interface{} to slice
	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		return
	}

	if len(slice) == 0 {
		return
	}

	linkedList.val = slice[0]

	pNode := new(LinkedNode)
	linkedList.next = pNode

	for i, value := range slice {
		if i == 0 {
			continue
		}
		pNode.val = value
		pNode.next = new(LinkedNode)
		pNode = pNode.next
	}

	return
}
