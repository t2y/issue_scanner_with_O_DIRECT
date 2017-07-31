package main

import (
	"bufio"
	"log"
	"os"

	"github.com/ncw/directio"
)

func Read(path string) {
	f, err := directio.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("error: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error: ", err)
	}

	return
}

func main() {
	log.Println("start")

	Read("test.txt")

	log.Println("end")
}
