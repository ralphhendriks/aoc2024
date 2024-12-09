package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	test   int64
	values []int64
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var equations []equation

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p1, p2, found := strings.Cut(scanner.Text(), ": ")
		if !found {
			panic("unrecognized pattern")
		}
		t, err := strconv.ParseInt(p1, 10, 64)
		if err != nil {
			panic("conversion failed")
		}

		var values []int64
		for _, n := range strings.Fields(p2) {
			v, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				panic("conversion failed")
			}
			values = append(values, v)
		}

		equations = append(equations, equation{t, values})
	}

	fmt.Printf("Answer part 1: %d\n", part1(equations))
	fmt.Printf("Answer part 2: %d\n", part2(equations))
}

func part1(equations []equation) int64 {
	var ans int64 = 0
	for _, e := range equations {
		if test(e.values[0], e.values[1:], e.test) {
			ans += e.test
		}
	}
	return ans
}

func test(head int64, tail []int64, testValue int64) bool {
	if len(tail) > 1 {
		return test(head+tail[0], tail[1:], testValue) || test(head*tail[0], tail[1:], testValue)
	}
	return head+tail[0] == testValue || head*tail[0] == testValue
}

func part2(equations []equation) int64 {
	var ans int64 = 0
	for _, e := range equations {
		if test2(e.values[0], e.values[1:], e.test) {
			ans += e.test
		}
	}
	return ans
}

func test2(head int64, tail []int64, testValue int64) bool {
	if len(tail) > 1 {
		return test2(head+tail[0], tail[1:], testValue) || test2(head*tail[0], tail[1:], testValue) || test2(concat(head, tail[0]), tail[1:], testValue)
	}
	return head+tail[0] == testValue || head*tail[0] == testValue || concat(head, tail[0]) == testValue
}

func concat(lhs int64, rhs int64) int64 {
	v, err := strconv.ParseInt(fmt.Sprintf("%d%d", lhs, rhs), 10, 64)
	if err != nil {
		panic("conversion failes")
	}
	return v
}
