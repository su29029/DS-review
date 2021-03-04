package main

import (
	"fmt"
)

func main() {
	fmt.Println("stack")
	stack := new(Stack)
	stack.Create([]int{1, 2, 3, 4})
	fmt.Println(stack.GetTopValue())
	stack.Push(5)
	fmt.Println(stack.GetTopValue())
	stack.Pop()
	fmt.Println(stack.GetTopValue())
	stack.Pop()
	fmt.Println(stack.GetTopValue())
	stack.Pop()
	fmt.Println(stack.GetTopValue())
	stack.Pop()
	fmt.Println(stack.GetTopValue())
}
