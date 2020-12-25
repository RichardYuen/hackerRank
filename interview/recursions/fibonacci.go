package main

import "fmt"

// Solution
func fibonacci(n int) int {
	if(n == 0 || n == 1){
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Test
func main(){
	fmt.Printf("%d\n", fibonacci(3))
}