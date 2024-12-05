package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	fst, snd int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var rules []rule
	var updates [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rules = append(rules, parseRule(scanner.Text()))
	}

	for scanner.Scan() {
		updates = append(updates, parseUpdate(scanner.Text()))
	}

	ans := 0
	for _, u := range updates {
		if isRightOrder(u, rules) {
			ans += u[len(u)/2]
		}
	}
	fmt.Printf("Answer part 1: %d\n", ans)

	ans = 0
	for _, u := range updates {
		if !isRightOrder(u, rules) {
			c := correct(u, rules)
			ans += c[len(c)/2]
		}
	}
	fmt.Printf("Answer part 2: %d\n", ans)
}

func parseRule(s string) rule {
	fields := strings.Split(s, "|")
	i, err := strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}
	j, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	return rule{fst: i, snd: j}
}

func parseUpdate(s string) []int {
	var pages []int
	for _, p := range strings.Split(s, ",") {
		page, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		pages = append(pages, page)
	}
	return pages
}

func isRightOrder(update []int, rules []rule) bool {
	for _, r := range rules {
		if !applyRule(update, r) {
			return false
		}
	}
	return true
}

func applyRule(update []int, r rule) bool {
	fst := slices.Index(update, r.fst)
	snd := slices.Index(update, r.snd)
	return fst < 0 || snd < 0 || fst < snd
}

// Think of bubble sort, but backwards.
// If the rules don't match then start swapping pairs from the back.
func correct(update []int, rules []rule) []int {
	c := slices.Clone(update)
	for i := 1; i < len(c); i++ {
		for j := i - 1; j >= 0; j-- {
			if isRightOrder(c[0:i+1], rules) {
				break
			}
			tmp := c[j]
			c[j] = c[j+1]
			c[j+1] = tmp
		}
	}
	return c
}
