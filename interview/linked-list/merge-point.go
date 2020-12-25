// Problem: Find Merge Point of Two Lists

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

// Abs returns the absolute value of x.
func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}

// Solution
func findMergeNode(headA *Node, headB *Node) int{
	var i, j *Node
	var lenA, lenB, d int

	lenA = 0
	lenB = 0

	i = headA
	for(i != nil){
		lenA++
		i = i.next
	}
	i = headB
	for(i != nil){
		lenB++
		i = i.next
	}

	d = Abs(lenA - lenB)
	i = headA
	j = headB
	if(lenA > lenB){
		for(d > 0){
			i = i.next
			d--
		}
	}else{
		for(d > 0){
			j = j.next
			d--
		}
	}
	for(i != j){
		i = i.next
		j = j.next
	}
	return i.data
}

//Test
func test1(){
	headA := makeNode(1)
	headA.next = makeNode(2)
	headB := makeNode(1)
	headB.next = headA.next
	headA.next = makeNode(3)
	fmt.Printf("%d\n", findMergeNode(headA, headB))
}

func test2(){
	headA := makeNode(1)
	headA.next = makeNode(2)
	headA.next.next = makeNode(3)
	headB := makeNode(1)
	headB.next = headA.next.next 
	fmt.Printf("%d\n", findMergeNode(headA, headB))
}

func main(){
	test1()
	test2()
}