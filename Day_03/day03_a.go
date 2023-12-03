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
	
	var sumOfParts int

	for i := 0; i<len(engine); i++{
		for j:=0; j<len(engine); j++{
			if unicode.IsDigit(engine[i][j]) &&
				(isSymbol(engine[i-1][j]) || isSymbol(engine[i-1][j-1]) || 
				isSymbol(engine[i][j-1]) || isSymbol(engine[i+1][j-1]) || 
				isSymbol(engine[i+1][j]) || isToSum(i,j+1, engine)){

				var toSum string
				for k := 0; unicode.IsDigit(engine[i][j+k]); k++{
					toSum += string(engine[i][j+k])
				}

				j += len(toSum)-1
				toSumtoInt, _ := strconv.Atoi(toSum)
				sumOfParts += toSumtoInt
			}
		}
	}
	fmt.Println(sumOfParts)
}

func isSymbol(r rune)bool{
	return r!='.' && !unicode.IsDigit(r)
}

func isToSum (i, j int, engine [][]rune)bool{
	if !unicode.IsDigit(engine[i][j]){
		j--
	}

	if unicode.IsDigit(engine[i][j+1]){
		return isSymbol(engine[i-1][j]) || isSymbol(engine[i+1][j]) || isToSum(i, j+1, engine)
	}

	return isSymbol(engine[i-1][j]) || 
	isSymbol(engine[i-1][j+1]) || 
	isSymbol(engine[i][j+1]) || 
	isSymbol(engine[i+1][j+1]) || 
	isSymbol(engine[i+1][j])
}