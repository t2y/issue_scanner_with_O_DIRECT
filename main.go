package main

import (
	"flag"
	"log"
	"os"
	"syscall"
)

var direct bool

func Read(path string) {
	var (
		f   *os.File
		err error
	)

	if direct {
		f, err = os.OpenFile(path, os.O_RDONLY|syscall.O_DIRECT, 0)
	} else {
		f, err = os.OpenFile(path, os.O_RDONLY, 0)
	}
	if err != nil {
		log.Println("error: ", err)
	}
	defer f.Close()

	data := make([]byte, 100)
	for {
		n, err := f.Read(data)
		log.Print(n, ", ", string(data))
		if err != nil {
			log.Println("err: ", err)
			break
		}
	}

	return
}

func main() {
	log.Println("start")

	flag.BoolVar(&direct, "direct", false, "set O_DIRECT")
	flag.Parse()

	Read("test.txt")

	log.Println("end")
}
