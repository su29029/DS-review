package main

import "fmt"

func main() {
	var linkedList = new(LinkedList)
	var arr = []int{1, 2, 3, 4, 5}
	linkedList.Create(arr)
	for node := linkedList; node.next != nil; node = node.next {
		fmt.Println(node.val)
	}
	return
}
