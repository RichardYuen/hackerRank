// Problem: Lowest Common Ancestor

package main

import "fmt"

//Auxillary code
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
func lca(root *Node, value1 int, value2 int) *Node{
	if(value1 < root.data && value2 < root.data){
		return lca(root.left, value1, value2)
	}
	if(value1 > root.data && value2 > root.data){
		return lca(root.right, value1, value2)
	}
	return root
}

// Test
func main(){
	root := makeNode(4)
	root.left = makeNode(2)
	root.right = makeNode(7)
	root.left.left = makeNode(1)
	root.left.right = makeNode(3)
	root.right.left = makeNode(6)
	fmt.Printf("%d\n", lca(root, 1, 7).data)
}