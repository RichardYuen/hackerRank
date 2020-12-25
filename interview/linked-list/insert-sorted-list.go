// Problem: Inserting a Node Into a Sorted Doubly Linked List

package main

import "fmt"

// Auxillary code
type Node struct{
	data int
	next *Node
	prev *Node
}

func makeNode(data int) *Node{
	var new Node
	new.data = data
	new.prev = nil
	new.next = nil
	return &new
}

func printList(head *Node){
	if(head == nil){
		fmt.Println()
	}else{
		fmt.Printf("%d ", head.data)
		printList(head.next)
	}
}

// Solution
func sortedInsert(head *Node, data int) *Node{
	var current, temp *Node
	temp = makeNode(data)

	if(head == nil){
		return temp
	}

	if(head.data >= data){
		head.prev = temp
		temp.next = head
		return temp
	}

	current = head
	for(current.next != nil && current.next.data < data){
		current = current.next
	}
	temp.next = current.next
	if(current.next != nil){
		current.next.prev = temp
	}
	current.next = temp
	temp.prev = current
	return head
}

// Test
func main(){
	head := makeNode(1)
	head.prev = nil 
	head.next = makeNode(3)
	head.next.prev = head
	head.next.next = makeNode(4)
	head.next.next.prev = head.next
	head.next.next.next = makeNode(10)
	head = sortedInsert(head, 5)
	printList(head)
}