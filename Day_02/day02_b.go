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

	for sc.Scan(){
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
		sum += extracted["red"] * extracted["green"] * extracted["blue"]
	}

	fmt.Println(sum)
}