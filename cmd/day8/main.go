package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
	instruction string
	addendum    int
}

func main() {
	ss := readInput()
	x := findPart1(ss)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(ss)
	fmt.Printf("Part 2: %d\n", x)
}

func readInput() []command {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return getCommands(strings.Split(string(b), "\n"))
}

func getCommands(ss []string) []command {
	var cc []command
	for i := range ss {
		s := strings.Split(ss[i], " ")
		x, _ := strconv.Atoi(s[1])
		cc = append(cc, command{
			instruction: s[0],
			addendum:    x,
		})
	}

	return cc
}

func execute(maxRun int, cc []command) int {
	acc := 0
	m := make(map[int]int)
	idx := 0
	stop := false
	for !stop {
		if m[idx] == maxRun {
			break
		}

		c := cc[idx]
		m[idx]++
		switch c.instruction {
		case "acc":
			acc += c.addendum
			idx++
		case "jmp":
			idx += c.addendum
		case "nop":
			idx++
		default:
			panic(fmt.Sprintf("invalid command: '%s'", c.instruction))
		}
	}

	return acc
}

func executeCorrectly(cc []command) (int, bool) {
	acc := 0
	m := make(map[int]int)
	idx := 0
	stop := false
	for !stop {
		if m[idx] == 1 {
			return acc, false
		}
		if idx >= len(cc) {
			return acc, true
		}
		c := cc[idx]
		m[idx]++
		switch c.instruction {
		case "acc":
			acc += c.addendum
			idx++
		case "jmp":
			idx += c.addendum
		case "nop":
			idx++
		default:
			panic(fmt.Sprintf("invalid command: '%s'", c.instruction))
		}
	}
	return -1, false
}

func findPart1(cc []command) int {
	return execute(1, cc)
}

func findPart2(cc []command) int {
	for i := range cc {
		t := cc[i]

		switch t.instruction {
		case "acc":
			continue
		case "jmp":
			cc[i].instruction = "nop"
		case "nop":
			cc[i].instruction = "jmp"
		}
		if a, b := executeCorrectly(cc); b {
			return a
		}

		cc[i] = t
	}

	return -1
}
