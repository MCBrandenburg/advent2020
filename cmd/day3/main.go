package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	ss := readInput("./input.txt")
	fmt.Println(ss[0])
	x := findPart1(ss)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(ss)
	fmt.Printf("Part 2: %d\n", x)
}
func readInput(f string) [][]string {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	ss := strings.Split(string(b), "\n")
	tt := make([][]string, len(ss))
	for i := range tt {
		tt[i] = make([]string, len(ss[0]))
	}
	for y, s := range ss {
		for x, r := range s {
			tt[y][x] = string(r)
		}
	}

	return tt
}

func findPart1(ss [][]string) int {
	return run(ss, 1, 3)
}

func findPart2(ss [][]string) int {
	runs := []struct {
		x int
		y int
	}{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	t := 1

	for _, r := range runs {
		t *= run(ss, r.y, r.x)
	}

	return t
}

func run(ss [][]string, slopeY int, slopeX int) int {
	count := 0
	x := slopeX
	xlen := len(ss[0])
	for y := slopeY; y < len(ss); y += slopeY {
		if ss[y][x] == "#" {
			count++
		}
		x += slopeX
		x = x % xlen
	}
	return count
}
