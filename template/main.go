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

type Input struct{}

func ReadInput(reader *bufio.Reader) *Input {
	return &Input{}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input := ReadInput(reader)
	ret := Solution(input)
	fmt.Fprintln(writer, ret)
}

func Solution(input *Input) any {
	return nil
}
