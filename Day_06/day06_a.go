package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)
	
	var times, distances []int
	
	sc.Scan()
	for _, value := range strings.Fields(sc.Text())[1:]{
		time, _ := strconv.Atoi(value)
		times = append(times, time)
	}
	sc.Scan()
	for _, value := range strings.Fields(sc.Text())[1:]{
		distance, _ := strconv.Atoi(value)
		distances = append(distances, distance)
	}

	var product int = 1

	for i, time := range times{
		var wayToWin int
		for releaseTime := 0; releaseTime<time; releaseTime++{
			if releaseTime*(time-releaseTime)>distances[i]{
				wayToWin++
			}
		}
		product *= wayToWin
	}
	fmt.Println(product)
}