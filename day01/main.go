package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Answer part 1: %d\n", part1(f))

	// rewind
	f.Seek(0, io.SeekStart)

	fmt.Printf("Answer part 2: %d\n", part2(f))
}

func part1(reader io.Reader) int {
	l, r, ans := []int{}, []int{}, 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		p, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		l = append(l, p)

		q, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		r = append(r, q)
	}

	slices.Sort(r)
	slices.Sort(l)

	for i := 0; i < len(l); i++ {
		if l[i] > r[i] {
			ans += l[i] - r[i]
		} else {
			ans += r[i] - l[i]
		}
	}
	return ans
}

func part2(reader io.Reader) int {
	l, r, ans := []int{}, make(map[int]int), 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		p, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		l = append(l, p)

		q, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		r[q] += 1
	}

	for _, v := range l {
		ans += v * r[v]
	}
	return ans
}
