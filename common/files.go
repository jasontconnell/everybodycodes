package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
