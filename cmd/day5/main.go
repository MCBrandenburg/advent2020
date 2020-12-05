package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	ss := readInput()
	x := findPart1(ss)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(ss, x)
	fmt.Printf("Part 2: %d\n", x)
}

func findPart1(ss []string) int {
	max := 0
	for _, s := range ss {
		if id := calcSeatID(s); id > max {
			max = id
		}
	}
	return max
}

func findPart2(ss []string, max int) int {
	m := make(map[int]struct{})

	for _, s := range ss {
		id := calcSeatID(s)
		m[id] = struct{}{}
	}

	id := -1
	for i := 0; i <= max; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		_, xok := m[i-1]
		_, yok := m[i+1]
		if xok && yok {
			id = i
		}
	}

	return id
}

func readInput() []string {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func calcSeatID(s string) int {
	var (
		row  int
		seat int
	)

	if s[0] == 'B' {
		row += 64
	}

	if s[1] == 'B' {
		row += 32
	}

	if s[2] == 'B' {
		row += 16
	}

	if s[3] == 'B' {
		row += 8
	}

	if s[4] == 'B' {
		row += 4
	}

	if s[5] == 'B' {
		row += 2
	}

	if s[6] == 'B' {
		row += 1
	}

	if s[7] == 'R' {
		seat += 4
	}
	if s[8] == 'R' {
		seat += 2
	}

	if s[9] == 'R' {
		seat += 1
	}

	return row*8 + seat
}
