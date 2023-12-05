package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
	"math"
)

type interval struct{
	start, end int
}

func (i interval) len()int{
	return i.end-i.start
}

func (i interval) apply(from, to interval)(res [3]interval){
	if from.start > i.start{
		res[0] = interval{i.start, from.start}
	}
	res[1]=interval{int(math.Max(float64(i.start), float64(from.start)))+(to.start-from.start), 
						int(math.Min(float64(i.end), float64(from.end)))+(to.end-from.end)}
	if from.end < i.end{
		res[2] = interval{from.end, i.end}
	}
	return
}

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)
	
	sc.Scan()
	var intervals []interval
	var startSeed int
	for i, seedString := range(strings.Fields(sc.Text())[1:]){
		if i%2==0{
			startSeed, _ = strconv.Atoi(seedString)
			continue
		}
		numberOfSeeds, _ := strconv.Atoi(seedString)
		intervals = append(intervals, interval{startSeed,startSeed + numberOfSeeds})
	}

	sc.Scan();sc.Scan()
		
	var newIntervals,leftIntervals []interval

	for sc.Scan(){
		if sc.Text()==""{
			newIntervals = append(newIntervals, intervals...)
			intervals = newIntervals
			newIntervals = make([]interval, 0)
			continue
		}
		var startTo, startFrom, length int
		fmt.Sscanf(sc.Text(), "%d %d %d", &startTo, &startFrom, &length)
		from := interval{startFrom, startFrom + length}
		to := interval{startTo, startTo+length}
		
		for _, interval := range intervals {
			if interval.end > from.start && interval.start < from.end{
				result := interval.apply(from, to)
				if result[0].len()>0{
					leftIntervals = append(leftIntervals, result[0])
				}
				if result[2].len()>0{
					leftIntervals = append(leftIntervals, result[2])
				}
				newIntervals = append(newIntervals, result[1])
				continue
			}
			leftIntervals = append(leftIntervals, interval)
		}
		intervals = leftIntervals
		leftIntervals = make([]interval, 0)
	}

	var minLocation int
	for i, interval := range append(newIntervals, intervals...){
		if interval.start < minLocation || i == 0{
			minLocation = interval.start
		}
	}
	fmt.Println(minLocation)
}
