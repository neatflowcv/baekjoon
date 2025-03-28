package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Read[T any](reader *bufio.Reader) T {
	var ret T
	_, err := fmt.Fscan(reader, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

type Input struct {
	n       int // 사람 수
	persons []*Person
}

type Person struct {
	no  int
	arr []int
}

func ReadInput(reader *bufio.Reader) *Input {
	n := Read[int](reader)
	var persons []*Person
	for i := 0; i < n; i++ {
		no := Read[int](reader)
		var arr []int
		for j := 0; j < 4; j++ {
			arr = append(arr, Read[int](reader))
		}
		persons = append(persons, &Person{
			no:  no,
			arr: arr,
		})
	}
	return &Input{
		n:       n,
		persons: persons,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input := ReadInput(reader)
	ret := Solution(input)
	fmt.Fprint(writer, ret[0])
	for i := 1; i < len(ret); i++ {
		fmt.Fprint(writer, " ", ret[i])
	}
	fmt.Fprintln(writer)
}

func Solution(input *Input) []int {
	sort.Slice(input.persons, func(i, j int) bool {
		return input.persons[i].no < input.persons[j].no
	})
	black := make(map[int]struct{})
	var ret []int
	for i := 0; i < 4; i++ {
		maxScore := -1
		maxNo := -1
		for _, person := range input.persons {
			if _, ok := black[person.no]; ok {
				continue
			}
			if maxScore < person.arr[i] {
				maxScore = person.arr[i]
				maxNo = person.no
			}
		}
		black[maxNo] = struct{}{}
		ret = append(ret, maxNo)
	}
	return ret
}
