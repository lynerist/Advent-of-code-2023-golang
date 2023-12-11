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
	for i:=0; sc.Scan(); i+=2{
		expanded :=" " +strings.Join(strings.Split(sc.Text(), ""), " ") + " "
		pipes = append(pipes, []rune(expanded))
		if strings.Contains(expanded, "S"){
			start.y = i
			start.x = strings.Index(expanded, "S")
			pipes[start.y][start.x] = '*'
		}
		pipes = append(pipes, []rune(strings.Repeat(" ", len(pipes[0]))))
	}

	var position point
	var direction string

	switch{
	case start.y>0 && strings.Contains("7|F", string(pipes[start.y-2][start.x])):
		position = point{start.x, start.y-2}
		pipes[start.y-1][start.x] = '*' 
		direction="N"
	case start.x>0 && strings.Contains("F-L", string(pipes[start.y][start.x-2])):
		position = point{start.x-2, start.y}
		pipes[start.y][start.x-1] = '*' 
		direction="W"
	case start.y<len(pipes)-1 && strings.Contains("J|L", string(pipes[start.y+2][start.x])):
		position = point{start.x, start.y+2}
		pipes[start.y+1][start.x] = '*' 
		direction="S"		
	case start.x<len(pipes[0])-1 && strings.Contains("J-7", string(pipes[start.y][start.x+2])):
		position = point{start.x+2, start.y}
		pipes[start.y][start.x+1] = '*' 
		direction="E"
	}
	
	whereToGo:= map[string]modifier{"|N":{0,-1,"N"}, "|S":{0,1,"S"}, "-E":{1,0,"E"}, "-W":{-1,0,"W"},
	"LS":{1,0,"E"}, "LW":{0,-1,"N"}, "JS":{-1,0,"W"}, "JE":{0,-1,"N"},
	"7N":{-1,0,"W"}, "7E":{0,1,"S"}, "FN":{1,0,"E"}, "FW":{0,1,"S"}}
	expand := map[string]rune{"N":'|', "S":'|', "W":'-', "E":'-'}
	
	for pipes[position.y][position.x]!='*'{

		toNext := whereToGo[string(pipes[position.y][position.x])+direction]
		toNextNext := whereToGo[string(expand[toNext.direction])+toNext.direction]
		pipes[position.y+toNext.dy][position.x+toNext.dx] = '*'
		pipes[position.y][position.x] = '*'

		position = point{position.x+toNext.dx+toNextNext.dx, position.y+toNext.dy+toNextNext.dy}
		direction = toNextNext.direction
	}

	pipes = append([][]rune{[]rune(strings.Repeat(" ", len(pipes[0])))},pipes...)

	cleanOutside(0,0,pipes)
	
	var contained int
	for _, l := range pipes{
		for _, r := range l{
			if r!='*' && r!= 0 && r!= ' '{
				contained++
			}
		}
	}
	fmt.Println(contained)
}

func cleanOutside(x,y int, pipes [][]rune){
	if y<0 || x<0 || y==len(pipes) || x==len(pipes[0]) || pipes[y][x] == 0 || pipes[y][x] == '*' {
		return
	}
	pipes[y][x] = 0
	
	cleanOutside(x+1, y, pipes)
	cleanOutside(x, y+1, pipes)
	cleanOutside(x-1, y, pipes)
	cleanOutside(x, y-1, pipes)
}
