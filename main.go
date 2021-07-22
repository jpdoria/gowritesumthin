package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

// Reads the lines from the file
func read() string {
	data, err := ioutil.ReadFile("the-matrix.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return string(data)
}

// Outputs the lines from the text file to stdout
func write(data string) {
	time.Sleep(10 * time.Second)
	for {
		robotgo.TypeStr(data, 500.0) // 500ms
		time.Sleep(1 * time.Second)
	}
}

func main() {
	data := read()
	write(data)
}
