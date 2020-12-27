// Problem: Is This a Binary Search Tree

package main

import "fmt"

// Auxillary code
type Node struct{
	data int
	left *Node
	right *Node
}

func newNode(data int) *Node{
	var new Node
	new.data = data
	new.left = nil
	new.right = nil
	return &new
}

// Solution
func bstUtil(root *Node, min int, max int) bool{
	if(root == nil){
		return true
	}
	if(root.data < min || root.data > max){
		return false
	}
	return bstUtil(root.left, min, root.data-1) && bstUtil(root.right, root.data+1, max)
}

func checkBST(root *Node) bool{
	return bstUtil(root, 0, 10000)
}

// Test
func main(){
	tree := newNode(4)
	tree.left = newNode(2)
	tree.right = newNode(6)
	tree.left.left = newNode(1)
	tree.left.right = newNode(3)
	tree.right.left = newNode(5)
	tree.right.right = newNode(7)
	if(checkBST(tree)){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}

}