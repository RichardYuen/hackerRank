// Problem: Crossword Puzzle

package main

import (
	"fmt"
	"strings"
)

// Auxillary code
func printSolution(solution []string){
	for i := 0; i < 10; i++ {
		fmt.Println(solution[i])
	}
}

func solveUtil(crossword []string) bool{


}

func crosswordPuzzle(crossword []string, words string) []string{

}

// Test  
func main(){
	crossword := []string{
		"+-++++++++",
		"+-++++++++",
		"+-++++++++",
		"+-----++++",
		"+-+++-++++",
		"+-+++-++++",
		"+++++-++++",
		"++------++",
		"+++++-++++",
		"+++++-++++",
	}

	words := "LONDON;DELHI;ICELAND;ANKARA"

	printSolution(crosswordPuzzle(crossword, words))
}