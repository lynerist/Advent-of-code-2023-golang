package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct{
	x, y int
}

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	universe := make(map[point]bool)
	
	var galaxyInX, galaxyInY []bool

	for y:=0; sc.Scan(); y++{
		if y==0{
			galaxyInX = make([]bool, len(sc.Text()))
		}
		galaxyInY = append(galaxyInY, false)
		
		for x, r := range sc.Text(){
			if r=='#'{
				universe[point{x,y}] = true
				galaxyInX[x] = true
				galaxyInY[y] = true
			}
		}
	}

	toMove := make(map[point]int)
	for x, isThereGalaxy := range galaxyInX{
		if !isThereGalaxy{
			for galaxy := range universe{
				if galaxy.x>x{
					toMove[galaxy]++
				}
			}
		}
	}

	var moved []point
	for galaxy, distance := range toMove{
		moved = append(moved, point{galaxy.x+distance*(1000000-1), galaxy.y})
		delete(universe, galaxy)
	}
	for _,galaxy := range moved{
		universe[galaxy] = true
	}

	toMove = make(map[point]int)
	for y, isThereGalaxy := range galaxyInY{
		if !isThereGalaxy{
			for galaxy := range universe{
				if galaxy.y>y{
					toMove[galaxy]++
				}
			}
		}
	}
	moved = make([]point, 0)
	for galaxy, distance := range toMove{
		moved = append(moved, point{galaxy.x, galaxy.y+distance*(1000000-1)})
		delete(universe, galaxy)
	}
	for _,galaxy := range moved{
		universe[galaxy] = true
	}

	var sumDistances int

	for galaxy := range universe{
		for otherGalaxy := range universe{
			sumDistances += int(math.Abs(float64((galaxy.x-otherGalaxy.x)))) +
							int(math.Abs(float64((galaxy.y-otherGalaxy.y))))
		}
	}
	fmt.Println(sumDistances/2)
}
