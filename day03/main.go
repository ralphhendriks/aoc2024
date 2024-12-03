package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer part 1: %d\n", part1(b))
	fmt.Printf("Answer part 2: %d\n", part2(b))
}

func part1(b []byte) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	ans := 0
	for _, m := range re.FindAllSubmatch(b, -1) {
		l, err := strconv.Atoi(string(m[1]))
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(string(m[2]))
		if err != nil {
			panic(err)
		}

		ans += l * r
	}
	return ans
}

func part2(b []byte) int {
	re := regexp.MustCompile(`(?:mul|don\'t|do)\((?:(\d+),(\d+))?\)`)
	reading, ans := true, 0
	for _, m := range re.FindAllSubmatch(b, -1) {
		s := string(m[0])
		switch {
		case s == "do()":
			reading = true
		case s == "don't()":
			reading = false
		case reading:
			l, err := strconv.Atoi(string(m[1]))
			if err != nil {
				panic(err)
			}

			r, err := strconv.Atoi(string(m[2]))
			if err != nil {
				panic(err)
			}
			ans += l * r
		}
	}
	return ans
}
