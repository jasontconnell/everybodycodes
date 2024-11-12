package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadBytes(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Println("can't read file", filename)
	}
	return b
}

func ReadInts(filename string) []int {
	lines := ReadLines(filename)
	if lines == nil {
		return nil
	}

	list := []int{}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Println("can't get int from", line)
			continue
		}
		list = append(list, i)
	}
	return list
}

func ReadIntGrid(filename string) [][]int {
	lines := ReadLines(filename)
	if lines == nil {
		return nil
	}

	list := make([][]int, len(lines))
	for y := 0; y < len(lines); y++ {
		f := strings.Fields(lines[y])
		list[y] = make([]int, len(f))
		for x := 0; x < len(f); x++ {
			val, err := strconv.Atoi(f[x])
			if err != nil {
				log.Println("invalid int", f[x])
				continue
			}
			list[y][x] = val
		}
	}
	return list
}

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("can't read file", filename)
		return nil
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}
	return lines
}
