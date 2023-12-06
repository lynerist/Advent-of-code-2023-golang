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
	
	sc.Scan()
	time, _ := strconv.Atoi(strings.Join(strings.Fields(sc.Text())[1:],""))
	sc.Scan()
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(sc.Text())[1:],""))

	var wayToWin int
	for releaseTime := 0; releaseTime<time; releaseTime++{
		if releaseTime*(time-releaseTime)>distance{
			wayToWin++
		}
	}
	fmt.Println(wayToWin)
}