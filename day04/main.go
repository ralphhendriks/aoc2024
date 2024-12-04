package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var p []string
	for scanner.Scan() {
		p = append(p, scanner.Text())
	}

	fmt.Printf("Answer part 1: %d\n", part1(p))
	fmt.Printf("Answer part 2: %d\n", part2(p))
}

func part1(p []string) int {
	ans := 0
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[0]); y++ {
			ans += checkAllXmas(p, x, y)
		}
	}
	return ans
}

func checkAllXmas(p []string, x int, y int) int {
	ans := checkXmas(p, x, y, 0, 1)
	ans += checkXmas(p, x, y, 1, 0)
	ans += checkXmas(p, x, y, 0, -1)
	ans += checkXmas(p, x, y, -1, 0)
	ans += checkXmas(p, x, y, 1, 1)
	ans += checkXmas(p, x, y, 1, -1)
	ans += checkXmas(p, x, y, -1, 1)
	ans += checkXmas(p, x, y, -1, -1)
	return ans
}

func checkXmas(p []string, x int, y int, dx int, dy int) int {
	const XMAS string = "XMAS"
	for idx := range XMAS {
		i := x + idx*dx
		j := y + idx*dy
		if i < 0 || i > len(p)-1 || j < 0 || j > len(p[0])-1 || p[i][j] != XMAS[idx] {
			return 0
		}
	}
	return 1
}

func part2(p []string) int {
	ans := 0
	for x := 1; x < len(p)-1; x++ {
		for y := 1; y < len(p[0])-1; y++ {
			ans += checkMas(p, x, y)
		}
	}
	return ans
}

func checkMas(p []string, x int, y int) int {
	if p[x][y] == 'A' && p[x-1][y-1] == 'M' && p[x+1][y+1] == 'S' && p[x-1][y+1] == 'M' && p[x+1][y-1] == 'S' {
		return 1
	}
	if p[x][y] == 'A' && p[x-1][y-1] == 'M' && p[x+1][y+1] == 'S' && p[x-1][y+1] == 'S' && p[x+1][y-1] == 'M' {
		return 1
	}
	if p[x][y] == 'A' && p[x-1][y-1] == 'S' && p[x+1][y+1] == 'M' && p[x-1][y+1] == 'M' && p[x+1][y-1] == 'S' {
		return 1
	}
	if p[x][y] == 'A' && p[x-1][y-1] == 'S' && p[x+1][y+1] == 'M' && p[x-1][y+1] == 'S' && p[x+1][y-1] == 'M' {
		return 1
	}
	return 0
}
