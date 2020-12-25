package main

import "fmt"

func minimumSwaps(arr []int) int{
	var n int
	var violates int


	n = len(arr)
	violates = 0

	for i:=1; i<=n; i++ {
		if(arr[i-1] != i){
			violates += 1
		}	
	}
	return violates-1
}

func main(){
	arr1 := []int{7,1,3,2,4,5,6}
	arr2 := []int{4,3,1,2}
	fmt.Printf("Minimum swap: %d\n", minimumSwaps(arr1[:]))
	fmt.Printf("Minimum swap: %d\n", minimumSwaps(arr2[:]))
}