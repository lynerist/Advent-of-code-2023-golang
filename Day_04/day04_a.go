package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var totalPoints int
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
		totalPoints += int(math.Pow(2, float64(myWinningNumbers-1)))
	}
	fmt.Println(totalPoints)
}