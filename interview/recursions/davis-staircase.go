// Problem: Davis' Staircase

package main

import (
    "fmt"
    "math"
)

// Auxillary code
var mem = initMem()

func initMem() []int{
    array := make([]int, 36)
    for i:=0;i<36;i++{
        array[i] = 0
    }
    return array
}

// Solution
func stepPerms(n int) int {
    var i int
    var result int
    var cnt int
    if(n == 0 || n == 1){
        return 1
    }
    
    cnt = 0
    for i=1;i<=3&&i<=n;i++{
        if(mem[n-i] != 0){
            result = mem[n-i]
        }else{
            result = stepPerms(n-i)
            mem[n-i] = result
        }
        cnt = cnt + result

    }
    return cnt % (int(math.Pow10(10))+7)
}

// Test
func main(){
    fmt.Printf("%d\n", stepPerms(1))
    fmt.Printf("%d\n", stepPerms(3))
    fmt.Printf("%d\n", stepPerms(7))
}