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
	ans := 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		x, err := sliceAtoi(fields)
		if err != nil {
			panic(err)
		}
		if safeIncreasing(x, false) || safeDecreasing(x, false) {
			ans += 1
		}
	}

	return ans
}

func sliceAtoi(v []string) ([]int, error) {
	var ret []int
	for _, s := range v {
		var i, err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}

func safeIncreasing(v []int, damp bool) bool {
	for i := 1; i < len(v); i++ {
		d := v[i] - v[i-1]
		if d < 1 || d > 3 {
			if damp {
				tmp := make([]int, len(v))
				copy(tmp, v)
				if safeIncreasing(slices.Delete(tmp, i-1, i), false) {
					return true
				}
				copy(tmp, v)
				return safeIncreasing(slices.Delete(tmp, i, i+1), false)
			}
			return false
		}
	}
	return true
}

func safeDecreasing(v []int, damp bool) bool {
	for i := 1; i < len(v); i++ {
		d := v[i] - v[i-1]
		if d < -3 || d > -1 {
			if damp {
				tmp := make([]int, len(v))
				copy(tmp, v)
				if safeDecreasing(slices.Delete(tmp, i-1, i), false) {
					return true
				}
				copy(tmp, v)
				return safeDecreasing(slices.Delete(tmp, i, i+1), false)
			}
			return false
		}
	}
	return true
}

func part2(reader io.Reader) int {
	ans := 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		x, err := sliceAtoi(fields)
		if err != nil {
			panic(err)
		}
		if safeIncreasing(x, true) || safeDecreasing(x, true) {
			ans += 1
		}
	}

	return ans
}
