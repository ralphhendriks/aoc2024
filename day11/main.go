package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var stones []int
	for _, s := range strings.Fields(string(b)) {
		if i, err := strconv.Atoi(s); err == nil {
			stones = append(stones, i)
		}
	}

	fmt.Printf("Answer part 1: %d\n", blink(stones, 25))
	fmt.Printf("Answer part 2: %d\n", blink(stones, 75))
}

func blink(stones []int, blinks int) int {
	hist := make(map[int]int)
	for _, s := range stones {
		hist[s] = 1
	}

	for i := 0; i < blinks; i++ {
		new_hist := make(map[int]int)
		for k, v := range hist {
			switch {
			case k == 0:
				new_hist[1] += v
			case digits(k)%2 == 0:
				lhs, rhs := split(k)
				new_hist[lhs] += v
				new_hist[rhs] += v
			default:
				new_hist[k*2024] += v
			}
		}
		hist = new_hist
	}

	ans := 0
	for _, v := range hist {
		ans += v
	}
	return ans
}

func digits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func split(i int) (int, int) {
	n := digits(i) / 2
	f := 10
	for k := 1; k < n; k++ {
		f *= 10
	}
	return i / f, i % f
}
