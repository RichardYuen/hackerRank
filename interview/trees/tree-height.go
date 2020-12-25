package main

import "fmt"

// Auxillary code 
type Node struct{
	data int
	left *Node
	right *Node
}

func makeNode(data int) *Node{
	var new Node
	new.data = data
	new.left = nil
	new.right = nil
	return &new
}

// Solution
func treeHeight(root *Node) int{
	var left, right int
	if(root == nil){
		return -1
	}else{
		left = treeHeight(root.left)
		right = treeHeight(root.right)
		if(left > right){
			return left + 1
		}else{
			return right + 1
		}
	}
}

// Test
func main(){
	var root *Node
	root = makeNode(3)
	root.left = makeNode(2)
	root.left.left = makeNode(1)
	root.right = makeNode(5)
	root.right.left = makeNode(4)
	root.right.right = makeNode(6)
	root.right.right.right = makeNode(7)
	fmt.Printf("%d\n", treeHeight(root))
}
