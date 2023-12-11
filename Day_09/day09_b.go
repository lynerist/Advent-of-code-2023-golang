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

		sum += previousInSequence(sequence)
	}
	fmt.Println(sum)

}

func previousInSequence(sequence []int)int{
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
	return sequence[0] - previousInSequence(nextSequence)
}
