package main

import (
	"fmt"
)

func main() {
	fmt.Println("Queue")
	queue := new(Queue)
	queue.Create([]int{1, 2, 3, 4, 5})
	fmt.Println(queue.head.val)
	fmt.Println(queue.tail.val)
	queue.Enqueue(6)
	fmt.Println(queue.tail.val)
	queue.Enqueue(7)
	fmt.Println(queue.tail.val)
	queue.Enqueue(8)
	fmt.Println(queue.tail.val)
	queue.Dequeue()
	fmt.Println(queue.head.val)
	queue.Dequeue()
	fmt.Println(queue.head.val)
	queue.Dequeue()
	fmt.Println(queue.head.val)
	queue.Dequeue()
	fmt.Println(queue.head.val)
	queue.Dequeue()
	fmt.Println(queue.head.val)
	queue.Dequeue()
	fmt.Println(queue.head.val)
}
