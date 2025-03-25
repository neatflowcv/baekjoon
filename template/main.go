package main

import (
	"bufio"
	"fmt"
	"os"
)

func Read[T any](reader *bufio.Reader) T {
	var ret T
	_, err := fmt.Fscan(reader, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func newReader() *bufio.Reader {
	_, ok := os.LookupEnv("DEBUG")
	if ok {
		file, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		return bufio.NewReader(file)
	}
	reader := bufio.NewReader(os.Stdin)
	return reader
}

type Input struct {
}

func ReadInput(reader *bufio.Reader) *Input {
	return &Input{}
}

func main() {
	reader := newReader()
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input := ReadInput(reader)
	_ = input
}
