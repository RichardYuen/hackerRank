// Problem: Huffman Decoding

package main

import "fmt"

// Auxillary code
type Node struct{
	data string
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
func decodeHuff(root *Node, s string){
	result := ""
	temp := root
	for i := 0; i < len(s); i++ {
		if(string(s[i]) == "0"){
			temp = temp.left
		}else{
			temp = temp.right
		}

		if(temp.left == nil && temp.right == nil){
			result += temp.data
			temp = root
		}
	}
	fmt.Println(result)
}
