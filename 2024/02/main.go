package main

import (
	"log"
	"strings"

	"github.com/jasontconnell/everybodycodes/common"
)

func main() {
	p1 := part1()
	p2 := part2()

	log.Println("Part 1:", p1)
	log.Println("Part 2:", p2)
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
		count += findSymbols(words, line)
	}
	return count
}

func findSymbols(words []string, line string) int {
	symbols := make(map[int]int)

	for _, w := range words {
		var idx int
		for {
			idx = indexStart(line, w, idx)
			if idx != -1 {
				for j := idx; j < idx+len(w); j++ {
					symbols[j] = 1
				}
				idx++
			} else {
				break
			}
		}

		rw := reverse(w)
		var ridx int
		for {
			ridx = indexStart(line, rw, ridx)
			if ridx != -1 {
				for j := ridx; j < ridx+len(rw); j++ {
					symbols[j] = 1
				}
				ridx++
			} else {
				break
			}
		}
	}

	return len(symbols)
}

func indexStart(s, substr string, start int) int {
	if start > len(s) {
		return -1
	}
	idx := strings.Index(s[start:], substr)
	if idx == -1 {
		return -1
	}
	return start + idx
}

func reverse(s string) string {
	var rw string
	for i := len(s) - 1; i >= 0; i-- {
		rw += string(s[i])
	}
	return rw
}
