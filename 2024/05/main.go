package main

import (
	"log"

	"github.com/jasontconnell/everybodycodes/common"
)

func main() {
	p1 := part1()

	log.Println("Part 1:", p1)
}

func part1() int {
	g := common.ReadIntGrid("everybody_codes_e2024_q5_p1.txt")
	list := wtf(g, 10)
	return list[len(list)-1]
}

func wtf(g [][]int, n int) []int {
	cp := rotate(g)
	cur := 0
	values := []int{}
	for i := 0; i < n; i++ {
		c := cur % len(cp)
		next := (cur + 1) % len(cp)
		top, ncpc := dequeue(cp[c])
		cp[c] = ncpc
		nlist := cp[next]
		dir := top / len(nlist)
		if dir%2 == 0 {
			nlist = insert(nlist, top%len(nlist)-1, top)
		} else {
			nlist = insert(nlist, len(nlist)-top%len(nlist), top)
		}
		cp[next] = nlist
		values = append(values, getNum(cp))
		cur++
	}
	return values
}

func getNum(g [][]int) int {
	val := 0
	mult := 1
	for y := len(g) - 1; y >= 0; y-- {
		val += g[y][0] * mult
		mult = mult * 10
	}
	return val
}

func dequeue(list []int) (int, []int) {
	x := list[0]
	list = list[1:]
	return x, list
}

func insert(list []int, idx, val int) []int {
	if idx < len(list) {
		suffix := append([]int{val}, list[idx:]...)
		list = append(list[:idx], suffix...)
	} else {
		list = append(list, val)
	}
	return list
}

func rotate(g [][]int) [][]int {
	ng := make([][]int, len(g[0]))
	for y := 0; y < len(g[0]); y++ {
		ng[y] = make([]int, len(g))
		for x := 0; x < len(g); x++ {
			ng[y][x] = g[x][y]
		}
	}
	return ng
}
