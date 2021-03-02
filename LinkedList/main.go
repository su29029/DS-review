package main

import "fmt"

func main() {
	var linkedList = new(LinkedList)
	var arr = []int{1, 2, 3, 4, 5}
	var err error

	err = linkedList.Create(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	linkedList.Print()
	linkedList.Add(6)
	linkedList.Print()
	err = linkedList.Modify(3, 7, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		linkedList.Print()
	}
	result, err := linkedList.Search(9, 1)
	fmt.Println(result)
	linkedList.Delete(6, 1)
	linkedList.Print()
	return
}
