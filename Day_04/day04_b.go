package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type specialQueue []int

func (s *specialQueue) add(n, copies int){
	for i:=0; i < len(*s) && i<n; i++{
		(*s)[i] += copies
	}
	for (n-len(*s) > 0){
		*s = append(*s, copies)
	}
}

func (s *specialQueue) pop()(n int){
	if len(*s)>0{
		n = (*s)[0]
		*s = (*s)[1:]
	}
	return
}

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var totalCards int
	var wonCards specialQueue
	for sc.Scan(){
		winningNumbers := make(map[string]bool)
		for _, number := range strings.Fields(strings.Split(strings.Split(sc.Text(), ":")[1], "|")[0]){
			winningNumbers[number] = true
		}
		var myWinningNumbers int
		for _, number := range strings.Fields(strings.Split(sc.Text(), "|")[1]){
			if winningNumbers[number] {
				myWinningNumbers++
			}
		}
		currentCards := 1 + wonCards.pop()
		totalCards += currentCards
		wonCards.add(myWinningNumbers, currentCards)
	}
	fmt.Println(totalCards)
}