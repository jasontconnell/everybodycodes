package main

import (
	"log"
	"strings"

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
	lines := common.ReadLines("everybody_codes_e2024_q2_p1.txt")
	words := getWords(lines)
	inst := findAllRunes(words, lines[2:])
	return inst
}

func part2() int {
	lines := common.ReadLines("everybody_codes_e2024_q2_p2.txt")
	words := getWords(lines)
	inst := findAllSymbols(words, lines[2:])
	return inst
}

func part3() int {
	lines := common.ReadLines("everybody_codes_e2024_q2_p3.txt")
	words := getWords(lines)
	return findGridSymbols(words, lines[2:])
}

func getWords(lines []string) []string {
	if len(lines) == 0 {
		return nil
	}

	if !strings.HasPrefix(lines[0], "WORDS:") {
		log.Println("invalid format", lines[0])
		return nil
	}

	words := strings.Split(strings.TrimPrefix(lines[0], "WORDS:"), ",")
	return words
}

func findAllRunes(words []string, lines []string) int {
	c := 0
	for _, w := range words {
		c += findRuneInstances(w, lines)
	}
	return c
}

func findRuneInstances(word string, text []string) int {
	count := 0

	for _, line := range text {
		tokens := strings.Fields(line)
		c := 0
		for _, t := range tokens {
			if strings.Contains(t, word) {
				c++
			}
		}
		count += c
	}
	return count
}

func findAllSymbols(words []string, lines []string) int {
	count := 0
	for _, line := range lines {
		s := findSymbols(words, line, false)
		count += len(s)
	}
	return count
}

func findSymbols(words []string, line string, wrap bool) []int {
	symbols := make(map[int]int)

	for _, w := range words {
		var idx int
		for {
			idx, _ = indexStart(line, w, idx, wrap)
			if idx != -1 {
				for j := idx; j < idx+len(w); j++ {
					c := j % len(line)
					symbols[c] = 1
				}
				idx++
				if idx >= len(line) {
					break
				}
			} else {
				break
			}
		}

		rw := reverse(w)
		var ridx int
		for {
			ridx, _ = indexStart(line, rw, ridx, wrap)
			if ridx != -1 {
				for j := ridx; j < ridx+len(rw); j++ {
					c := j % len(line)
					symbols[c] = 1
				}
				ridx++
				if ridx >= len(line) {
					break
				}
			} else {
				break
			}
		}
	}

	list := []int{}
	for v := range symbols {
		list = append(list, v)
	}
	return list
}

type coord struct {
	x, y int
}

func findGridSymbols(words []string, grid []string) int {
	m := make(map[coord]int)
	for i := 0; i < len(grid); i++ {
		list := findSymbols(words, grid[i], true)
		for _, v := range list {
			c := coord{x: v, y: i}
			m[c] = 1
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		cline := getVerticalLine(i, grid)
		list := findSymbols(words, cline, false)
		for _, v := range list {
			c := coord{x: i, y: v}
			m[c] = 1
		}
	}
	return len(m)
}

func getVerticalLine(col int, lines []string) string {
	s := ""
	for y := 0; y < len(lines); y++ {
		s += string(lines[y][col])
	}
	return s
}

func indexStart(s, substr string, start int, wrap bool) (int, bool) {
	if start > len(s) || len(substr) > len(s) {
		return -1, false
	}
	idx := strings.Index(s[start:], substr)
	if idx == -1 && !wrap {
		return -1, false
	} else {
		ns := s + s[:len(substr)]
		idx = strings.Index(ns[start:], substr)
		if idx == -1 {
			return -1, false
		}
	}
	return start + idx, wrap
}

func reverse(s string) string {
	var rw string
	for i := len(s) - 1; i >= 0; i-- {
		rw += string(s[i])
	}
	return rw
}
