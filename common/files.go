package common

import (
	"bufio"
	"log"
	"os"
)

func ReadBytes(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Println("can't read file", filename)
	}
	return b
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
