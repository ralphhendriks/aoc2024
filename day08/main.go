package main

import (
	"bufio"
	"fmt"
	"os"
)

type vector struct {
	x, y int
}

var antennas map[rune][]vector
var p, q int // max x and y coords

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	y := 0
	antennas = make(map[rune][]vector)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p = len(scanner.Text()) // todo: not very efficient to do this every line
		for x, c := range scanner.Text() {
			if c == '.' {
				continue
			}
			antennas[c] = append(antennas[c], vector{x, y})
		}
		y += 1
	}
	q = y
	fmt.Printf("Answer part 1: %d\n", allAntinodes(antinodes))
	fmt.Printf("Answer part 2: %d\n", allAntinodes(antinodes2))
}

func allAntinodes(f func(vector, vector) []vector) int {
	antinodes := make(map[vector]bool)
	for _, a := range antennas {
		for i := 0; i < len(a)-1; i++ { // from the first to the one before last antenna
			for j := i + 1; j < len(a); j++ { // from the i+1-th to the last antenna
				for _, x := range f(a[i], a[j]) {
					antinodes[x] = true
				}
			}
		}
	}
	return len(antinodes)
}

func antinodes(a vector, b vector) []vector {
	var rc []vector
	an1, an2 := vector{2*a.x - b.x, 2*a.y - b.y}, vector{2*b.x - a.x, 2*b.y - a.y}
	if inGrid(an1) {
		rc = append(rc, an1)
	}
	if inGrid(an2) {
		rc = append(rc, an2)
	}
	return rc
}

func antinodes2(a vector, b vector) []vector {
	var rc []vector
	dx := b.x - a.x
	dy := b.y - a.y
	for n := a; inGrid(n); n = (vector{n.x - dx, n.y - dy}) {
		rc = append(rc, n)
	}
	for n := b; inGrid(n); n = (vector{n.x + dx, n.y + dy}) {
		rc = append(rc, n)
	}
	return rc
}

func inGrid(v vector) bool {
	return v.x >= 0 && v.y >= 0 && v.x < p && v.y < q
}
