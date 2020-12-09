package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	ss := readInput()
	fmt.Println(ss[0])
	x := findPart1(ss)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(ss, x)
	fmt.Printf("Part 2: %d\n", x)
}

func readInput() []int {
	var ii []int

	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	ss := strings.Split(string(b), "\n")
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		ii = append(ii, i)
	}
	return ii
}

func findPart1(input []int) int {
	return findPart1Preabmle(25, input)
}

func findPart1Preabmle(preamble int, input []int) int {
	for i := preamble; i < len(input); i++ {
		s := input[i]
		if findPair(s, input[i-preamble:i]) == -1 {
			return s
		}
	}

	return -1
}

func findPair(s int, ii []int) int {
	m := make(map[int]struct{})
	for _, i := range ii {
		m[i] = struct{}{}
		x := s - i
		if _, ok := m[x]; ok {
			return x * i
		}
	}
	return -1
}

func findPart2(input []int, target int) int {
	for i := 0; i < len(input); i++ {
		min := input[i]
		max := input[i]
		t := input[i]
		for j := i + 1; j < len(input); j++ {
			t += input[j]
			if t > target {
				break
			}
			if input[j] < min {
				min = input[j]
			}
			if input[j] > max {
				max = input[j]
			}
			if t == target {
				return min + max
			}
		}
	}

	return -1
}
