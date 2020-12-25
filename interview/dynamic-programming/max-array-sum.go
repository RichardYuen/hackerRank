// Problem: Max Array Sum 

package main

import "fmt"

// Auxillary code
func max(a int, b int) int{
	if(a > b){
		return a
	}
	return b
}

// Solution
func maxSubsetSum(array []int) int{
	var current_max int

	n := len(array)
	memory := make([]int, n)

	if(n == 1){
		return array[0]
	}else if(n == 2){
		return max(array[0], array[1])
	}

	memory[0] = array[0]
	memory[1] = max(array[0], array[1])

	current_max = max(memory[0], memory[1])
	for i := 2; i < n; i++ {
		memory[i] = max(array[i], max(current_max, array[i]+memory[i-2]))
		current_max = memory[i]
	}

	return memory[n-1]
}

// Test
func main(){
	arr1 := []int{3,7,4,6,5}
	arr2 := []int{2,1,5,8,4}
	arr3 := []int{3,5,-7,8,10}
	fmt.Printf("%d\n", maxSubsetSum(arr1))
	fmt.Printf("%d\n", maxSubsetSum(arr2))
	fmt.Printf("%d\n", maxSubsetSum(arr3))	
}