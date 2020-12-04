package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type rulePart1 struct {
	min int
	max int
	r   rune
}

func (r rulePart1) IsValid(s string) bool {
	m := make(map[rune]int)
	_ = strings.Map(func(r rune) rune {
		m[r]++
		return r
	}, s)
	if m[r.r] >= r.min && m[r.r] <= r.max {
		return true
	}
	return false
}

func createRulePart1(s string) (rulePart1, string) {
	s = strings.Map(func(r rune) rune {
		switch r {
		case ' ', '-':
			return ','
		case ':':
			return -1
		default:
			return r
		}
	}, s)

	ss := strings.Split(s, ",")
	min, _ := strconv.Atoi(ss[0])
	max, _ := strconv.Atoi(ss[1])
	b := ss[2][0]
	return rulePart1{min: min, max: max, r: rune(b)}, ss[3]
}

func main() {

	ss := readInput()

	x := findPart1(ss)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(ss)
	fmt.Printf("Part 2: %d\n", x)
}

func findPart1(ss []string) int {
	count := 0
	for i := range ss {
		r, s := createRulePart1(ss[i])
		if r.IsValid(s) {
			count++
		}
	}
	return count
}

func findPart2(ss []string) int {
	count := 0
	for i := range ss {
		r, s := createRulePart2(ss[i])
		if r.IsValid(s) {
			count++
		}
	}
	return count
}

func readInput() []string {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	ss := strings.Split(string(b), "\n")
	return ss
}

type rulePart2 struct {
	pos1 int
	pos2 int
	r    rune
}

func (r rulePart2) IsValid(s string) bool {
	m := make(map[rune]int)
	_ = strings.Map(func(r rune) rune {
		m[r]++
		return r
	}, s)
	p1 := rune(s[r.pos1])
	p2 := rune(s[r.pos2])

	switch {
	case p1 == r.r && p2 == r.r:
		return false
	case p1 == r.r || p2 == r.r:
		return true
	default:
		return false
	}
}

func createRulePart2(s string) (rulePart2, string) {
	s = strings.Map(func(r rune) rune {
		switch r {
		case ' ', '-':
			return ','
		case ':':
			return -1
		default:
			return r
		}
	}, s)

	ss := strings.Split(s, ",")
	min, _ := strconv.Atoi(ss[0])
	max, _ := strconv.Atoi(ss[1])
	b := ss[2][0]
	return rulePart2{pos1: min - 1, pos2: max - 1, r: rune(b)}, ss[3]
}
