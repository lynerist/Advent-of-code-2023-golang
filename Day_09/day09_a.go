package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var sum int

	for sc.Scan(){
		var sequence []int
		for _, value := range strings.Fields(sc.Text()){
			intValue, _ := strconv.Atoi(value)
			sequence = append(sequence, intValue)
		}

		sum += nextInSequence(sequence)
	}
	fmt.Println(sum)

}

func nextInSequence(sequence []int)int{
	allZeroes := true
	var nextSequence []int
	for i, n := range sequence{
		allZeroes = allZeroes && n==0
		if i<len(sequence)-1{
			nextSequence = append(nextSequence, sequence[i+1]-n)
		}
	}
	if allZeroes{
		return 0
	}
	return sequence[len(sequence)-1] + nextInSequence(nextSequence)
}