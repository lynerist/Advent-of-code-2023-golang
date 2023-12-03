package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var engine [][]rune
	engine = append(engine, []rune{})
	for sc.Scan(){
		engine = append(engine, []rune(".." + sc.Text() + ".."))
	}
	engine = append(engine, []rune(strings.Repeat(".", len(engine[1]))))
	engine[0] = []rune(strings.Repeat(".", len(engine[1])))
	

	var sumOfRatios int
	possibleGears := make(map[string][]int)

	for i := 0; i<len(engine); i++{
		for j:=0; j<len(engine); j++{
			alreadyAdded := make(map[string]bool)

			if unicode.IsDigit(engine[i][j]) {
				var partNumber string
				for k := 0; unicode.IsDigit(engine[i][j+k]); k++{
					partNumber += string(engine[i][j+k])
				}

				leftIndices := [][]int{{-1,0},{-1,-1},{0,-1},{1,-1},{1,0}}
				for _, indices := range leftIndices{
					if isGear(engine[i+indices[0]][j+indices[1]]){
						addPossibleGear(i+indices[0], j+indices[1], possibleGears, partNumber, alreadyAdded)
					}
				}

				checkForGears(i, j+1, engine, possibleGears, partNumber, alreadyAdded)

				if len(partNumber)> 0{
					j += len(partNumber)-1
				}
			}
		}
	}
	for _, partNumbers := range possibleGears{
		if len(partNumbers) == 2{
			sumOfRatios += partNumbers[0] * partNumbers[1]
		}
	}
	fmt.Println(sumOfRatios)
}

func addPossibleGear(i, j int, possibleGears map[string][]int, partNumber string, alreadyAdded map[string]bool){
	gear := fmt.Sprintf("%d %d", i, j)
	partNumberToInt, _ := strconv.Atoi(partNumber)
	if !alreadyAdded[gear]{
		possibleGears[gear] = append(possibleGears[gear], partNumberToInt)
	}
	alreadyAdded[gear] = true
}

func isGear(r rune)bool{
	return r == '*'
}

func checkForGears (i, j int, engine [][]rune, possibleGears map[string][]int, partNumber string, alreadyAdded map[string]bool ){
	if !unicode.IsDigit(engine[i][j]){
		j--
	}

	if unicode.IsDigit(engine[i][j+1]){
		if isGear(engine[i-1][j]){
			addPossibleGear(i-1, j, possibleGears, partNumber, alreadyAdded)
		}
		if isGear(engine[i+1][j]){
			addPossibleGear(i+1, j, possibleGears, partNumber, alreadyAdded)
		}
		checkForGears(i, j+1, engine, possibleGears, partNumber, alreadyAdded)
	}

	rightIndices := [][]int{{-1,+1},{0,+1},{1,1}}
	if len(partNumber) > 1{
		rightIndices= append(rightIndices, [][]int{{-1,0}, {1,0}}...)
	}
	for _, indices := range rightIndices{
		if isGear(engine[i+indices[0]][j+indices[1]]){
			addPossibleGear(i+indices[0], j+indices[1], possibleGears, partNumber, alreadyAdded)
		}
	}
}