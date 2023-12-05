package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var seeds []int
	sc.Scan()
	for _, seedString := range(strings.Fields(sc.Text())[1:]){
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}
	sc.Scan(); sc.Scan()

	almanac := []func(int)int{}

	var maps []string
	for sc.Scan(){
		if sc.Text()==""{
			almanac = append(almanac, factory(maps))
			maps = make([]string, 0)
			sc.Scan()
			continue
		}
		maps = append(maps, sc.Text())
	}
	almanac = append(almanac, factory(maps))

	var minLocation int
	for i, seed := range seeds{
		location := seed
		for _, mapping := range almanac{
			location = mapping(location)
		}
		if location < minLocation || i==0{
			minLocation = location
		}
	}
	fmt.Println(minLocation)
}

func factory(maps []string) func(int)int{
	return func(n int)int{
		for _, mapLine := range maps{
			var start, startShifted, length int
			fmt.Sscanf(mapLine, "%d %d %d", &startShifted, &start, &length)
			if n>start && n<=start+length{
				return startShifted+n-start
			}
		}
		return n
	}
}