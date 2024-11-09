package main

import (
	"log"

	"github.com/jasontconnell/everybodycodes/common"
)

func main() {
	p1input := common.ReadBytes("everybody_codes_e2024_q1_p1.txt")
	p2input := common.ReadBytes("everybody_codes_e2024_q1_p2.txt")
	p3input := common.ReadBytes("everybody_codes_e2024_q1_p3.txt")

	log.Println("Part 1:", calculate(p1input, 1))
	log.Println("Part 2:", calculate(p2input, 2))
	log.Println("Part 3:", calculate(p3input, 3))
}

func getPotions(b byte) int {
	switch b {
	case 'A':
		return 0
	case 'B':
		return 1
	case 'C':
		return 3
	case 'D':
		return 5
	}
	return 0
}

func calculate(input []byte, teamsize int) int {
	total := 0
	for i := 0; i < len(input); i += teamsize {
		tcount := 0
		for j := 0; j < teamsize; j++ {
			p := getPotions(input[i+j])
			total += p
			if input[i+j] != 'x' {
				tcount++
			}
		}
		if tcount > 1 {
			total += tcount * (tcount - 1)
		}
	}
	return total
}
