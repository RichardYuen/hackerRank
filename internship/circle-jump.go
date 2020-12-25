// Problem: Circle Jumping

package main

import "fmt"

func findPath(nodes []int){
	n := len(nodes)
	visited := make([]bool, n)
	
	for i:=0;i<n;i++{
		visited[i] = false
	}

	for i:=0;i<n;i++{
		fmt.Printf("Start: node (%d) at index %d\n", nodes[i], i)
		result := traversal(nodes, visited, i, n)
		if(result){
			fmt.Println("Completed...")
			return
		}
		fmt.Println("Blocked!")
	}

	fmt.Println("Path not found!")
}

func traversal(nodes []int,visited []bool,pos int, len int)bool{
	if(visited[pos]){
		return false
	}

	visited[pos] = true
	if(checkAllVisited(visited)){
		return true
	}

	leftPos := ((pos - nodes[pos]) % len + len)%len
	rightPos := (pos + nodes[pos]) % len
	resultLeft := traversal(nodes, visited, leftPos, len)
	resultRight := traversal(nodes, visited, rightPos, len)
	if(!resultLeft && !resultRight){
		visited[pos] = false
		return false
	}else{
		if(resultLeft){
			fmt.Printf("Left from node(%d) at %d to node(%d) at %d\n", nodes[pos], pos, nodes[leftPos], leftPos)
		} else {
			fmt.Printf("Right from node(%d) at %d to node(%d) at %d\n", nodes[pos], pos, nodes[rightPos], rightPos)
		}
		return true
	}
}

func checkAllVisited(visited []bool)bool{
	for i:=0;i<len(visited);i++{
		if(!visited[i]){
			return false
		}
	}
	return true
}

func main(){
	findPath([]int{1,2,6,5,4,8,3,4,5,7})
	findPath([]int{2,6,1})
	findPath([]int{1,2,4,1,3})
	findPath([]int{1,2,3})
}