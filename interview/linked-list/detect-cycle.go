// Problem: Linked Lists: Detect a Cycle

package main

import "fmt"

// Auxillary code
type Node struct{
	data int 
	next *Node
}

func makeNode(data int) *Node{
	var new Node
	new.data = data
	new.next = nil
	return &new
}

// Solution
func hasCycle(head *Node) bool{
	var p1, p2 *Node

	p1 = head
	p2 = head
	for(p1 != nil && p2 != nil && p2.next != nil){
		p1 = p1.next
		p2 = p2.next.next
		if(p1 == p2){
			return true
		}
	}
	return false
}

// Test
func main(){
	head1 := makeNode(1)
	head2 := makeNode(1)
	head2.next = makeNode(2)
	head2.next.next = head2.next
	fmt.Printf("%t\n%t\n", hasCycle(head1), hasCycle(head2))
}