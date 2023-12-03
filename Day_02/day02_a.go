package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var sum int
	const red = 12
	const green = 13
	const blue = 14

	for i:=1; sc.Scan(); i++ {
		extracted := make(map[string]int)
		for _, sets := range strings.Split(strings.Split(sc.Text(), ":")[1], "; "){
			for _, set := range strings.Split(sets, ", "){
				var number int
				var color string
				fmt.Sscanf(set, "%d %s", &number, &color)
				if number > extracted[color] {
					extracted[color] = number
				}
			}
		}
		fmt.Println(extracted)
		if (extracted["red"]<=red && extracted["green"]<=green && extracted["blue"]<=blue){
			sum += i
		}
	}

	fmt.Println(sum)
}