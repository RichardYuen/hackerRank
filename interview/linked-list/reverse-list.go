// Problem: Reverse a doubly linked list

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
func reverse(head *Node) *Node{
	var temp *Node
	if(head == nil){
		return nil
	}
	if(head.next == nil){
		head.next = head.prev
		head.prev = nil
		return head
	}
	temp = head.prev
	head.prev = head.next
	head.next = temp
	return reverse(head.prev)
}

// Test
func main(){
	head := makeNode(1)
	head.prev = nil
	head.next = makeNode(2)
	head.next.prev = head
	head.next.next = makeNode(3)
	head.next.next.prev = head.next
	head.next.next.next = makeNode(4)
	head.next.next.next.prev = head.next.next
	head = reverse(head)
	printList(head)
}