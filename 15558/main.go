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

type Input struct {
	n     int
	k     int
	left  string
	right string
}

func ReadInput(reader *bufio.Reader) *Input {
	return &Input{
		n:     Read[int](reader),
		k:     Read[int](reader),
		left:  Read[string](reader),
		right: Read[string](reader),
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input := ReadInput(reader)
	ret := Solution(input)
	switch ret {
	case true:
		fmt.Fprintln(writer, "1")
	case false:
		fmt.Fprintln(writer, "0")
	}
}

type Location struct {
	line int // 0 or 1
	pos  int // 0 ~ n-1
}

func (l *Location) Next(k int) []*Location {
	return []*Location{
		{line: l.line, pos: l.pos + 1},
		{line: l.line, pos: l.pos - 1},
		{line: (l.line + 1) % 2, pos: l.pos + k},
	}
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.line, l.pos)
}

type Lanes struct {
	lanes []*Lane
}

func NewLanes(lanes []*Lane) *Lanes {
	return &Lanes{
		lanes: lanes,
	}
}

func (l *Lanes) Tick() {
	for _, lane := range l.lanes {
		lane.Tick()
	}
}

func (l *Lanes) Seconds() int {
	return l.lanes[0].seconds
}

func (l *Lanes) IsAvailable(line int, pos int) bool {
	return l.lanes[line].IsAvailable(pos)
}

type Lane struct {
	arr     []bool // true: 유저가 이동할 수 있다.
	seconds int    // 1초마다 발판이 사라진다.
}

func NewLane(arr string) *Lane {
	var available []bool
	for _, v := range arr {
		available = append(available, v == '1')
	}
	return &Lane{
		arr:     available,
		seconds: 0,
	}
}

func (l *Lane) Tick() {
	l.seconds++
}

func (l *Lane) IsAvailable(pos int) bool {
	if pos < 0 || len(l.arr) <= pos {
		return l.seconds < pos
	}
	return l.arr[pos] && l.seconds < pos
}

// 유저의 승리 가능 여부를 반환한다.
func Solution(input *Input) bool {
	user := &Location{
		line: 0,
		pos:  0,
	}
	lanes := NewLanes([]*Lane{
		NewLane(input.left),
		NewLane(input.right),
	})

	var currentLocations []*Location
	currentLocations = append(currentLocations, user)

	for len(currentLocations) > 0 && lanes.Seconds() < input.n {
		var nextCandidates []*Location
		for _, currentLocation := range currentLocations {
			nextCandidates = append(nextCandidates, currentLocation.Next(input.k)...)
		}
		var availableCandidates []*Location
		unique := make(map[string]struct{})
		for _, candidate := range nextCandidates {
			if _, ok := unique[candidate.String()]; ok {
				continue
			}
			unique[candidate.String()] = struct{}{}
			if lanes.IsAvailable(candidate.line, candidate.pos) {
				if candidate.pos >= input.n-1 {
					return true
				}
				availableCandidates = append(availableCandidates, candidate)
			}
		}
		lanes.Tick()
		currentLocations = availableCandidates
	}
	return false
}
