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
	var starts []string
	for sc.Scan(){
		var node, left, right string
		fmt.Sscanf(sc.Text(), "%3s = (%3s, %3s)", &node, &left, &right)
		nodes[node] = map[rune]string{'L':left, 'R':right}
		if node[2] == 'A'{
			starts = append(starts, node)
		}
	}

	passByZEvery := make(map[string]int)

	for i, currentNode := range starts{
		for steps := 0; ; steps++{
			if currentNode[2] == 'Z'{
				passByZEvery[starts[i]] = steps
				break
			}
			currentNode = nodes[currentNode][instructions[steps%len(instructions)]]
		}
	}

	var mcm = passByZEvery["AAA"]
	for found := false; !found; {
		found = true
		for _, coincidence := range passByZEvery{
			if mcm%coincidence!=0{
				found = false
				break
			}
		}
		mcm += passByZEvery["AAA"]
	}

	fmt.Println(mcm)
}

//15726453850399