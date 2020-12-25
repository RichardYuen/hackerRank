// Problem: Balanced Brackets

package main

import "fmt"

// Auxillary code
type Node struct{
	data string
	next *Node
}

type Stack struct{
	top *Node
}

func newNode(data string) *Node{
	var new Node
	new.data = data
	new.next = nil
	return &new
}

func (s *Stack)push(data string){
	node := newNode(data)
	node.next = s.top
	s.top = node
}

func (s *Stack)pop(){
	if(s.top == nil){
		return
	}
	s.top = s.top.next
}

func (s *Stack)isEmpty() bool{
	return s.top == nil
}

// Solution
func isBalanced(s string) string{
	var stack Stack
	for i := 0; i < len(s); i++ {
		if(stack.isEmpty()){
			stack.push(string(s[i]))
		}else{
			if(stack.top.data == "{"  && string(s[i]) == "}"){
				stack.pop()
			}else if(stack.top.data == "(" && string(s[i]) == ")"){
				stack.pop()
			}else if(stack.top.data == "[" && string(s[i]) == "]"){
				stack.pop()
			}else {
				stack.push(string(s[i]))
			}
		}
	}
	if(stack.isEmpty()){
		return "YES"
	}else{
		return "NO"
	}
}

// Test
func main(){
	fmt.Printf("%s\n", isBalanced("}][}}(}][))]"))
	fmt.Printf("%s\n", isBalanced("[](){()}"))
	fmt.Printf("%s\n", isBalanced("()"))
	fmt.Printf("%s\n", isBalanced("({}([][]))[]()"))
	fmt.Printf("%s\n", isBalanced("{)[](}]}]}))}(())("))
	fmt.Printf("%s\n", isBalanced("([[)"))
}