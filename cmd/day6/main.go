package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type group map[rune]int

func main() {
	ss := readInput("./input.txt")
	x := findPart1(ss)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(ss)
	fmt.Printf("Part 2: %d\n", x)
}

func readInput(f string) []group {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var (
		pp []group
		s  string
		c  int
	)

	for scanner.Scan() {
		t := scanner.Text()

		switch {
		case t == "":
			p := creatGroup(s, c)
			pp = append(pp, p)
			s = ""
			c = 0
		default:
			s += t
			c++
		}
	}
	p := creatGroup(s, c)
	pp = append(pp, p)

	return pp
}

func creatGroup(s string, c int) group {
	g := make(group)
	_ = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			g[r]++
		}
		return r
	}, s)
	g[0] = c
	return g
}

func findPart1(gg []group) int {
	sum := 0
	for _, g := range gg {
		sum += len(g) - 1
	}
	return sum
}

func findPart2(gg []group) int {
	sum := 0
	for _, g := range gg {
		want := g[0]
		for k, v := range g {
			if k == 0 {
				continue
			}
			if v == want {
				sum++
			}
		}
	}
	return sum
}
