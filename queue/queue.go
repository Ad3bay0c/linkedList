package queue

import "fmt"

type QueueNode struct {
	next	*QueueNode
	value	int
}

type QueueList struct {
	head	*QueueNode
}

func (q *QueueList) Enqueue(val int) {
	currentNode := q.head
	newNode := &QueueNode{value: val}
	if q.head == nil {
		q.head = newNode
		return
	}
	for currentNode != nil {
		if currentNode.next == nil {
			currentNode.next = newNode
			return
		}
		currentNode = currentNode.next
	}
}

func (q *QueueList) Dequeue() {
	currentNode := q.head
	if q.head != nil {
		q.head = currentNode.next
	}

}

func (q QueueList) Print() {
	currentNode := q.head
	for currentNode != nil {
		fmt.Print(currentNode.value," ")
		currentNode = currentNode.next
	}
}

func QueueImplementation() {
	queue := QueueList{}

	queue.Enqueue(23)
	queue.Enqueue(25)
	queue.Enqueue(35)
	queue.Print()
	fmt.Println()
	fmt.Println("........Dequeue.........")
	queue.Dequeue()
	queue.Print()
}