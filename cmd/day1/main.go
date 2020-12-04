package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	i := readInput()
	x := find2020Pair(i)
	fmt.Printf("Part 1: %d\n", x)
	x = find2020Triplet(i)
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

func find2020Pair(ii []int) int {
	m := make(map[int]struct{})
	for _, i := range ii {
		m[i] = struct{}{}
		x := 2020 - i
		if _, ok := m[x]; ok {
			fmt.Printf("(%d,%d)\n", i, x)
			return x * i
		}
	}
	return 0
}

func find2020Triplet(ii []int) int {
	for _, x := range ii {
		for _, y := range ii {
			for _, z := range ii {
				if (x + y + z) == 2020 {
					fmt.Printf("(%d,%d,%d)\n", x, y, z)
					return x * y * z
				}
			}
		}
	}

	return 0
}
