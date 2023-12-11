package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct{
	x, y int
}

type modifier struct{
	dx,dy int
	direction string
}

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var pipes [][]rune
	var start point
	for i:=0; sc.Scan(); i++{
		pipes = append(pipes, []rune(sc.Text()))
		if strings.Contains(sc.Text(), "S"){
			start.y = i
			start.x = strings.Index(sc.Text(), "S")
		}
	}

	var paths []point
	var directions []string

	switch{
	case start.y>0 && strings.Contains("7|F", string(pipes[start.y-1][start.x])):
		paths = append(paths, point{start.x, start.y-1})
		directions = append(directions, "N")
		fallthrough
	case start.x>0 && strings.Contains("F-L", string(pipes[start.y][start.x-1])):
		paths = append(paths, point{start.x-1, start.y})
		directions = append(directions, "W")
		fallthrough
	case start.y<len(pipes)-1 && strings.Contains("J|L", string(pipes[start.y+1][start.x])):
		paths = append(paths, point{start.x, start.y+1})
		directions = append(directions, "S")
		fallthrough
	case start.x<len(pipes[0])-1 && strings.Contains("J-7", string(pipes[start.y][start.x+1])):
		paths = append(paths, point{start.x+1, start.y})
		directions = append(directions, "E")
	}
	whereToGo:= map[string]modifier{"|N":{0,-1,"N"}, "|S":{0,1,"S"}, "-E":{1,0,"E"}, "-W":{-1,0,"W"},
	"LS":{1,0,"E"}, "LW":{0,-1,"N"}, "JS":{-1,0,"W"}, "JE":{0,-1,"N"},
	"7N":{-1,0,"W"}, "7E":{0,1,"S"}, "FN":{1,0,"E"}, "FW":{0,1,"S"}}

	var distance int 
	for distance = 1; paths[0] != paths[1]; distance++{
		for i:=0; i<2; i++{
			next := whereToGo[string(pipes[paths[i].y][paths[i].x])+directions[i]]
			paths[i] = point{paths[i].x+next.dx,paths[i].y+next.dy}
			directions[i] = next.direction
		}
	}
	fmt.Println(distance)
}
