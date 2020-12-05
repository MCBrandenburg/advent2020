package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

func main() {
	pp := readInput("./input.txt")
	x := findPart1(pp)
	fmt.Printf("Part 1: %d\n", x)
	x = findPart2(pp)
	fmt.Printf("Part 2: %d\n", x)
}

func readInput(f string) []passport {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var (
		pp []passport
		s  string
	)

	for scanner.Scan() {
		t := scanner.Text()

		switch {
		case t == "":
			p := createPassport(s)
			pp = append(pp, p)
			s = ""
		case s == "":
			s += t
		default:
			s += " " + t
		}
	}
	p := createPassport(s)
	pp = append(pp, p)

	return pp
}

func createPassport(s string) passport {
	ss := strings.Split(s, " ")
	p := make(passport)
	for i := range ss {
		sss := strings.Split(ss[i], ":")
		p[sss[0]] = sss[1]
	}

	return p
}

func findPart1(pp []passport) int {
	count := 0
	for i := range pp {
		if pp[i].IsValid() {
			count++
		}
	}
	return count
}

func findPart2(pp []passport) int {
	count := 0
	for i := range pp {
		if pp[i].IsValid2() {
			count++
		}
	}
	return count
}

func (p passport) IsValid() bool {
	_, byrOk := p["byr"]
	_, iyrOk := p["iyr"]
	_, eyrOk := p["eyr"]
	_, hgtOk := p["hgt"]
	_, hclOk := p["hcl"]
	_, eclOk := p["ecl"]
	_, pidOk := p["pid"]
	// _, cidOk := p["cid"]

	return byrOk && iyrOk && eyrOk && hgtOk && hclOk && eclOk && pidOk
}

func (p passport) IsValid2() bool {
	byrVal, byrOk := p["byr"]
	iyrVal, iyrOk := p["iyr"]
	eyrVal, eyrOk := p["eyr"]
	hgtVal, hgtOk := p["hgt"]
	hclVal, hclOk := p["hcl"]
	eclVal, eclOk := p["ecl"]
	pidVal, pidOk := p["pid"]
	// _, cidOk := p["cid"]

	if !(byrOk && iyrOk && eyrOk && hgtOk && hclOk && eclOk && pidOk) {
		return false
	}

	byrOk = validateByr(byrVal)
	iyrOk = validateIyr(iyrVal)
	eyrOk = validateEyr(eyrVal)
	hgtOk = validateHgt(hgtVal)
	hclOk = validateHcl(hclVal)
	eclOk = validateEcl(eclVal)
	pidOk = validatePid(pidVal)

	return byrOk && iyrOk && eyrOk && hgtOk && hclOk && eclOk && pidOk
}

func validateByr(val string) bool {
	i, err := strconv.Atoi(val)
	switch {
	case err != nil:
		return false
	case i >= 1920 && i <= 2002:
		return true
	default:
		return false
	}
}

func validateIyr(val string) bool {
	i, err := strconv.Atoi(val)
	switch {
	case err != nil:
		return false
	case i >= 2010 && i <= 2020:
		return true
	default:
		return false
	}
}

func validateEyr(val string) bool {
	i, err := strconv.Atoi(val)
	switch {
	case err != nil:
		return false
	case i >= 2020 && i <= 2030:
		return true
	default:
		return false
	}
}

func validateHgt(val string) bool {
	switch {
	case strings.Contains(val, "cm"):
		return validateCm(val)
	case strings.Contains(val, "in"):
		return validateIn(val)
	default:
		return false
	}
}

func validateCm(val string) bool {
	val = strings.ReplaceAll(val, "cm", "")
	i, err := strconv.Atoi(val)
	switch {
	case err != nil:
		return false
	case i >= 150 && i <= 193:
		return true
	default:
		return false
	}
}

func validateIn(val string) bool {
	val = strings.ReplaceAll(val, "in", "")
	i, err := strconv.Atoi(val)
	switch {
	case err != nil:
		return false
	case i >= 59 && i <= 76:
		return true
	default:
		return false
	}
}

func validateHcl(val string) bool {
	ok, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, val)
	return ok
}

func validateEcl(val string) bool {
	ok, _ := regexp.MatchString(`^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$`, val)
	return ok
}

func validatePid(val string) bool {
	ok, _ := regexp.MatchString(`^\d{9}$`, val)
	return ok
}
