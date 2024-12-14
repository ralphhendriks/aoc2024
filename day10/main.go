package main

import (
	"bufio"
	"fmt"
	"os"
)

type vector struct {
	x, y int
}

var trailmap [][]int

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	y := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := make([]int, len(scanner.Text()))
		for x, r := range scanner.Text() {
			l[x] = int(r - '0')
		}
		trailmap = append(trailmap, l)
		y++
	}

	ans1, ans2 := 0, 0
	for y := 0; y < len(trailmap); y++ {
		for x := 0; x < len(trailmap[0]); x++ {
			p := vector{x, y}
			ans1 += score(p, countUnique)
			ans2 += score(p, func(e []vector) int { return len(e) })
		}
	}
	fmt.Printf("Answer part 1: %d\n", ans1)
	fmt.Printf("Answer part 2: %d\n", ans2)
}

// returns the height at vector{x, y} or -1 if not on map
func height(p vector) int {
	if p.x < 0 || p.x > len(trailmap[0])-1 || p.y < 0 || p.y > len(trailmap)-1 {
		return -1
	}
	return trailmap[p.y][p.x]
}

func trails(p vector) []vector {
	h := height(p)
	if h == 9 {
		return []vector{p}
	}
	var next []vector
	if n := (vector{p.x, p.y - 1}); height(n) == h+1 {
		next = append(next, trails(n)...)
	}
	if s := (vector{p.x, p.y + 1}); height(s) == h+1 {
		next = append(next, trails(s)...)
	}
	if w := (vector{p.x + 1, p.y}); height(w) == h+1 {
		next = append(next, trails(w)...)
	}
	if e := (vector{p.x - 1, p.y}); height(e) == h+1 {
		next = append(next, trails(e)...)
	}
	return next
}

func score(p vector, agg func([]vector) int) int {
	if height(p) != 0 { // point is not a trailhead
		return 0
	}
	return agg(trails(p))
}

func countUnique(endpoints []vector) int {
	visited := make(map[vector]bool)
	for _, v := range endpoints {
		visited[v] = true
	}
	return len(visited)
}
