package main

import (
	"log"

	"github.com/jasontconnell/everybodycodes/common"
)

type xy struct {
	x, y int
}

func main() {
	p1 := part1()
	p2 := part2()
	p3 := part3()

	log.Println("Part 1:", p1)
	log.Println("Part 2:", p2)
	log.Println("Part 3:", p3)
}

func part1() int {
	lines := common.ReadLines("everybody_codes_e2024_q3_p1.txt")
	g := getGrid(lines)
	b := 0
	done := false
	total := 0
	level := 0
	w, h := len(lines[0]), len(lines)
	for !done {
		g, b = dig(w, h, g, level)
		total += b
		level++
		done = b == 0
	}
	return total
}
func part2() int {
	lines := common.ReadLines("everybody_codes_e2024_q3_p2.txt")
	g := getGrid(lines)
	b := 0
	done := false
	total := 0
	level := 0
	w, h := len(lines[0]), len(lines)
	for !done {
		g, b = dig(w, h, g, level)
		total += b
		level++
		done = b == 0
	}
	return total
}
func part3() int {
	return 0
}

func dig(w, h int, grid map[xy]int, level int) (map[xy]int, int) {
	m := make(map[xy]int)
	removed := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pt := xy{x, y}
			if b, ok := grid[pt]; ok {
				all := true
				for _, chkpt := range []xy{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
					if _, ok := grid[chkpt]; !ok || b != level {
						all = false
					}
				}
				if all || (level == 0 && b == 0) {
					m[pt] = level + 1
					removed++
				}
			}
		}
	}
	return m, removed
}

func getGrid(lines []string) map[xy]int {
	g := make(map[xy]int)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				g[xy{x, y}] = 0
			} else {
				g[xy{x, y}] = -1
			}
		}
	}
	return g
}
