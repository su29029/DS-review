package main

import (
	"fmt"
	"reflect"
	"sync"
)

// Stack the Stack data structure
type Stack struct {
	top    *node
	length int
	lock   *sync.RWMutex
}

type node struct {
	val  interface{}
	prev *node
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

// check if a slice
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

// Create create a stack
func (stack *Stack) Create(arr interface{}) error {
	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		return fmt.Errorf("not a slice")
	}

	stack.length = len(slice)
	stack.lock = &sync.RWMutex{}
	stack.top = new(node)

	if stack.length > 0 {
		for i, j := 0, stack.length-1; i <= j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}

		pNode := stack.top

		for i, value := range slice {
			pNode.val = value

			if i < len(slice)-1 {
				pNode.prev = new(node)
				pNode = pNode.prev
			}
		}
	}

	return nil
}

// GetTopValue return the top value of the stack
func (stack *Stack) GetTopValue() interface{} {
	if stack.length == 0 {
		return nil
	}
	return stack.top.val
}

// Push push a new node into the stack(on the top)
func (stack *Stack) Push(val interface{}) error {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	pNode := new(node)
	pNode.val = val
	pNode.prev = stack.top
	stack.top = pNode
	stack.length++
	return nil
}

// Pop pop the top node of the stack
func (stack *Stack) Pop() (ok bool, err error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.top == nil {
		err = fmt.Errorf("pop a empty stack")
		return
	}
	stack.top = stack.top.prev
	stack.length--
	ok = true
	return
}
