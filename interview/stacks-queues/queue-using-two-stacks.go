// Pro

package main 

import "fmt"

type Node struct{
	data int
	next *Node
}

type Stack struct{
	top *Node
}

type Queue struct{
	head *Stack
	tail *Stack
}

func newNode(data int) *Node{
	var new Node
	new.data = data
	new.next = nil
	return &new
}

func (s *Stack)push(data int){
	node := newNode(data)
	node.next = s.top
	s.top =  node
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

func (q *Queue)peek() int{
	if(q.head.isEmpty()){
		for(!q.tail.isEmpty()){
			data := q.tail.top.data
			q.tail.pop()
			q.head.push(data)
		}
	}
	return q.head.top.data
}

func (q *Queue)pop() int{
	if(q.head.isEmpty()){
		for(!q.tail.isEmpty()){
			data := q.tail.top.data
			q.tail.pop()
			q.head.push(data)
		}
	}
	result := q.head.top.data
	q.head.pop()
	return result
}

func (q *Queue)put(data int){
	q.tail.push(data)
}

func main(){
	var t, op, val int
	var queue Queue
	
	queue.head = &Stack{top: nil}
	queue.tail = &Stack{top: nil}

	fmt.Scanf("%d\n", &t)
	for i:=0; i<t; t++{
		fmt.Scanln(&op, &val)
		if(op == 1){
			queue.put(val)
		}else if (op == 2) {
			queue.pop()
		}else if (op == 3){
			fmt.Printf("%d\n", queue.peek())
		}
	}
}