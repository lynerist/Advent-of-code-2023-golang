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

	var sumAllCalibrationValues int
	for sc.Scan(){
		trimmed := strings.TrimFunc(sc.Text(), unicode.IsLetter)
		calibrationValue, _ := strconv.Atoi(trimmed[:1] + trimmed[len(trimmed)-1:])
		sumAllCalibrationValues += calibrationValue
	}
	fmt.Println(sumAllCalibrationValues)
}