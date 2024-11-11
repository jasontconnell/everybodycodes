package main

import (
	"log"
	"math"
	"sort"

	"github.com/jasontconnell/everybodycodes/common"
)

func main() {
	p1 := part1()
	p2 := part2()
	p3 := part3()

	log.Println("Part 1:", p1)
	log.Println("Part 2:", p2)
	log.Println("Part 3:", p3)
}

func part1() int {
	list := common.ReadInts("everybody_codes_e2024_q4_p1.txt")
	n := hammer(list)
	return n
}

func part2() int {
	list := common.ReadInts("everybody_codes_e2024_q4_p2.txt")
	n := hammer(list)
	return n
}

func part3() int {
	list := common.ReadInts("everybody_codes_e2024_q4_p3.txt")
	n := medianHammer(list)
	return n
}

func hammer(list []int) int {
	low := min(list)
	total := 0
	for _, v := range list {
		total += (v - low)
	}
	return total
}

func medianHammer(list []int) int {
	med := median(list)
	total := 0
	for _, v := range list {
		if v < med {
			total += med - v
		} else {
			total += v - med
		}
	}
	return total

}

func min(list []int) int {
	val := math.MaxInt32
	for _, v := range list {
		if v < val {
			val = v
		}
	}
	return val
}

func median(list []int) int {
	cp := make([]int, len(list))
	copy(cp, list)
	sort.Ints(cp)
	return cp[len(cp)/2]
}
