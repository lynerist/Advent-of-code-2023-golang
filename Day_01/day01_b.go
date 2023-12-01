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

	wordToDigit := map[string]string{"zero":"0","one":"1", "two":"2", "three":"3", "four":"4",
									"five":"5", "six":"6", "seven":"7", "eight":"8", "nine":"9"}

	var sumAllCalibrationValues int
	for sc.Scan(){
		var calibrationValueString, firstWord, lastWord string

		firstDigitIndex := strings.IndexFunc(sc.Text(), unicode.IsDigit)
		lastDigitIndex := strings.LastIndexFunc(sc.Text(), unicode.IsDigit)

		firstWordIndex := len(sc.Text())
		lastWordIndex := 0

		for word := range wordToDigit {
			firstIndexCurrentWord := strings.Index(sc.Text()[:firstDigitIndex], word)
			if firstIndexCurrentWord > -1 && firstIndexCurrentWord < firstWordIndex {
				firstWordIndex = firstIndexCurrentWord
				firstWord = word
			}
			lastIndexCurrentWord := strings.LastIndex(sc.Text()[lastDigitIndex:], word)
			if lastIndexCurrentWord > -1 && lastIndexCurrentWord > lastWordIndex {
				lastWordIndex = lastIndexCurrentWord
				lastWord = word
			}
		}
		if firstWord != "" {
			calibrationValueString += wordToDigit[firstWord]
		} else{
			calibrationValueString += sc.Text()[firstDigitIndex:firstDigitIndex+1]
		}

		if lastWord != "" {
			calibrationValueString += wordToDigit[lastWord]
		} else{
			calibrationValueString += sc.Text()[lastDigitIndex:lastDigitIndex+1]
		}

		calibrationValue, _ := strconv.Atoi(calibrationValueString)
		sumAllCalibrationValues += calibrationValue

	}
	fmt.Println(sumAllCalibrationValues)
}