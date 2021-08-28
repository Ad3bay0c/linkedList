package stack

import "fmt"

type Stack struct {
	items	[]interface{}
}

func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() interface{} {
	res := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]
	return res
}

func (s Stack) Peek() interface{} {
	return s.items[len(s.items)-1]
}

func (s Stack) PrintItems()  {
	fmt.Println(s.items)
}

func StackImplementation() {
	myStack := Stack{}
	myStack.Push("okay")
	myStack.Push("Yes")
	myStack.Push("Good")

	myStack.PrintItems()
	fmt.Println(myStack.Pop())
	myStack.PrintItems()

}
