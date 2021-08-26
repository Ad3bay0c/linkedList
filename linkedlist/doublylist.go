package linkedlist

import "fmt"

type DoublyNode struct {
	value	interface{}
	next	*DoublyNode
	prev	*DoublyNode
}

type DoublyLinkedList struct {
	head	*DoublyNode
	tail	*DoublyNode
	length	int
}

func (l DoublyLinkedList) PrintList() {
	if l.length == 0 {
		fmt.Println("No Data in the List")
		return
	}
	currentHead := l.head
	for i:=0; i<l.length;i++ {
		fmt.Printf("%d: %v\n", i+1, currentHead.value)
		currentHead = currentHead.next
	}
}

func (l *DoublyLinkedList) Add(value interface{}) {
	newNode := &DoublyNode{value: value}

	currentNode := l.head
	//previousNode := l.tail

	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}

	for currentNode !=nil {
		if currentNode.next == nil {
			newNode.prev = currentNode
			currentNode.next = newNode
			l.tail = newNode
			l.length++
			return
		}
		currentNode = currentNode.next
	}
}

func (l *DoublyLinkedList) Prepend(value interface{}) {
	newNode := &DoublyNode{value: value}
	currentNode := l.head

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}
	newNode.next = currentNode
	currentNode.prev = newNode
	l.head = newNode
	l.length++
}
func TestingDoubly() {
	list := DoublyLinkedList{}

	list.Add("First Ever Item")
	list.Add("Second Item")
	list.Add(32)

	list.PrintList()

	fmt.Println(".........Prepend.........")

	list.Prepend("Append to the first")
	list.Prepend("Append to the First Again")
	list.PrintList()

	fmt.Println(".........Search Item.........")
	fmt.Printf("The Value %v found at position: %d\n", 32, list.Search(32))

	fmt.Printf("The Value %v found at position: %d\n", 32, list.Search(23))

	fmt.Printf("The Value \"%v\" found at position: %v\n",  list.SearchAt(1), 1)
	fmt.Printf("The Value %v found at position: %d\n", list.SearchAt(5), 5)

	fmt.Println(".........Delete At.........")

	list.DeleteAt(1)
	list.DeleteAt(1)
	list.PrintList()

	fmt.Println(".........Insert Before.........")

	list.InsertBefore(1, "Insert before first element")
	list.InsertBefore(4, "Insert Before the last element")
	list.InsertBefore(4, "Insert Before Second Item")
	list.PrintList()
}

func (l DoublyLinkedList) Search(value interface{}) int {
	currentNode := l.head

	for i:=1; i<=l.length; i++ {
		if currentNode.value == value {
			return i
		}
		currentNode = currentNode.next
	}
	return -1
}

func (l DoublyLinkedList) SearchAt(pos int) interface{} {
	currentNode := l.head

	for i := 1; i<=l.length; i++ {
		if i == pos {
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return -1
}

func (l *DoublyLinkedList) DeleteAt(pos int) {
	currentNode := l.head
	lastNode := l.tail

	if pos < 1 || pos > l.length {
		fmt.Println("Invalid Position")
		return
	}
	if l.head == nil {
		fmt.Println("An Empty List")
		return
	}
	if pos == l.length {
		lastNode.prev.next = nil
		l.tail = lastNode.prev
		l.length--
		return
	}else if pos == 1 {
		currentNode.next.prev = nil
		l.head = currentNode.next
		l.length--
		return
	}
	for i:= 2; i<l.length; i++ {
		if i == pos {
			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			l.length--
			return
		}
		currentNode = currentNode.next
	}
}

func (l *DoublyLinkedList) InsertBefore(pos int, value interface{}) {
	newNode := &DoublyNode{value: value}
	currentNode := l.head
	lastNode := l.tail
	if pos > l.length || pos < 1 {
		fmt.Println("Index Out of range")
		return
	}
	if pos == 1 {
		newNode.next = currentNode
		currentNode.prev = newNode
		l.head = newNode
		l.length++
		return
	} else if pos == l.length{
		newNode.next = lastNode
		newNode.prev = lastNode.prev
		lastNode.prev.next = newNode
		lastNode.prev = newNode
		l.length++
		return
	}

	for i:=2; i<=l.length; i++ {
		if i == pos{
			newNode.prev = currentNode.prev
			newNode.next = currentNode
			currentNode.prev.next = newNode
			currentNode.prev = newNode
			l.length++
			return
		}
		currentNode = currentNode.next
	}

}