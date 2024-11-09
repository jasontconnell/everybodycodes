package common

import (
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
