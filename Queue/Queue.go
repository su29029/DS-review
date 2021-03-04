package main

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"unsafe"
)

// Queue the queue structure
type Queue struct {
	head *node
	tail *node
}

type node struct {
	val  interface{}
	next *node
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

// Create create a queue.
func (queue *Queue) Create(arr interface{}) error {
	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		return fmt.Errorf("not a slice")
	}

	queue.head = new(node)
	pNode := queue.head

	for i, val := range slice {
		pNode.val = val

		if i < len(slice)-1 {
			pNode.next = new(node)
			pNode = pNode.next
		}
	}

	queue.tail = pNode
	return nil
}

// Enqueue enqueue a node into the queue
func (queue *Queue) Enqueue(val interface{}) error {
	pNode := new(node)
	pNode.val = val
	pNode.next = nil

	tail := new(node)
	next := new(node)

	for {
		tail = queue.tail
		next = tail.next

		if tail != queue.tail {
			continue
		}

		if next != nil {
			// CAS(queue.tail, tail, next)
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&queue.tail)), unsafe.Pointer(tail), unsafe.Pointer(next))
			continue
		}

		// CAS(tail.next, next, pNode) == true
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next)), unsafe.Pointer(next), unsafe.Pointer(pNode)) {
			break
		}
	}

	// CAS(queue.tail, tail, pNode)
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&queue.tail)), unsafe.Pointer(tail), unsafe.Pointer(pNode))
	return nil
}

// Dequeue dequeue the first node of the queue
func (queue *Queue) Dequeue() error {

	head := new(node)
	tail := new(node)
	next := new(node)

	for {
		head = queue.head
		tail = queue.tail
		next = head.next

		if head != queue.head {
			continue
		}

		if head == tail && next == nil {
			return fmt.Errorf("dequeue an empty queue")
		}

		if head == tail && next == nil {

			// CAS(queue.tail, tail, next)
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&queue.tail)), unsafe.Pointer(tail), unsafe.Pointer(next))
			continue
		}

		// CAS(queue.head, head, next)
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&queue.head)), unsafe.Pointer(head), unsafe.Pointer(next)) {
			break
		}
	}

	return nil
}
