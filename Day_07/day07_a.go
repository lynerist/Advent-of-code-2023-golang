package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct{
	cards string
	bid int
	rank int
}

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)
	
	cardValue := map[rune]int{'T':10, 'J':11, 'Q':12, 'K':13, 'A':14}
	for r := '2'; r<='9'; r++{
		cardValue[r] = int(r-'0')
	}

	var hands []hand
	for sc.Scan(){
		input := strings.Fields(sc.Text())
		bid, _ := strconv.Atoi(input[1])
		byCard := make(map[rune]int)
		for _, r := range input[0]{
			byCard[r]++
		}
		
		var maxCopies int
		for _, copies := range byCard{
			if copies>maxCopies{
				maxCopies = copies
			}
		}
		var rank int
		switch maxCopies{
		case 5: rank = 6
		case 4: rank = 5
		case 3: rank = 3 + (3 - len(byCard))
		case 2: rank = 1 + (4 - len(byCard))
		}
		hands = append(hands, hand{input[0],bid, rank})
	}

	sort.Slice(hands, func(i,j int)bool{
		if hands[i].rank == hands[j].rank{
			for k := 0; k<len(hands[i].cards); k++{
				if cardValue[rune(hands[i].cards[k])]==cardValue[rune(hands[j].cards[k])]{
					continue
				}
				return cardValue[rune(hands[i].cards[k])]<cardValue[rune(hands[j].cards[k])]
			}
		}
		return hands[i].rank < hands[j].rank
	})

	var totalWinnings int
	for i, hand := range hands{
		totalWinnings += (i+1) * hand.bid
	}
	fmt.Println(totalWinnings)
}