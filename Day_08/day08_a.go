package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	sc.Scan()
	instructions := []rune(sc.Text())
	sc.Scan()
	
	nodes := make(map[string]map[rune]string)
	for sc.Scan(){
		var node, left, right string
		fmt.Sscanf(sc.Text(), "%3s = (%3s, %3s)", &node, &left, &right)
		nodes[node] = map[rune]string{'L':left, 'R':right}
	}

	var steps int
	currentNode := "AAA"
	for steps = 0; currentNode != "ZZZ"; steps++{
		currentNode = nodes[currentNode][instructions[steps%len(instructions)]]
	}

	fmt.Println(steps)
}