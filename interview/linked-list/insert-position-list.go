// Problem: Insert a node at a specific position in a linked list

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
func insertNth(head *Node, data int, position int) *Node{
	var i, j *Node

	n := makeNode(data)
	if(head == nil || position == 0){
		if(head != nil){
			n.next = head
		}
		head = n
		return head
	}

	i = head
	j = nil
	for (position > 0){
		j = i
		i = i.next
		position--
	}

	n.next = i
	j.next = n
	return head
}

// Test
func main(){
	head := makeNode(16)
	head.next = makeNode(13)
	head.next.next = makeNode(7)
	head = insertNth(head, 1, 2)
	printList(head)
}