package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type vector struct {
	x, y int
}

type step struct {
	x, y, dx, dy int
}

var lab []string
var start vector

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lab = append(lab, scanner.Text())
	}
	start = findStart()

	fmt.Printf("Answer part 1: %d\n", part1())
	fmt.Printf("Answer part 2: %d\n", part2())
}

func part1() int {
	return countUnique(path())
}

func findStart() vector {
	for y, l := range lab {
		x := strings.Index(l, "^")
		if x >= 0 {
			return vector{x, y}
		}
	}
	panic("start not found")
}

func path() []vector {
	l, d := start, vector{0, -1}

	var path []vector
	path = append(path, vector{l.x, l.y})

	for {
		next := add(l, d)

		if !inLab(next) {
			return path
		}

		if lab[next.y][next.x] == '#' {
			d = rotateRight(d)
			continue
		}

		l = next

		path = append(path, vector{l.x, l.y})
	}
}

func inLab(p vector) bool {
	return p.y >= 0 && p.y < len(lab) && p.x >= 0 && p.x < len(lab[0])
}

func rotateRight(v vector) vector {
	return vector{-v.y, v.x}
}

func add(v vector, w vector) vector {
	return vector{v.x + w.x, v.y + w.y}
}

func countUnique(path []vector) int {
	visited := make(map[vector]bool)
	for _, v := range path {
		visited[v] = true
	}
	return len(visited)
}

func part2() int {
	path := path()
	var obstructions []vector

	for i := 1; i < len(path); i++ {
		if checkLoop(path[i]) {
			obstructions = append(obstructions, path[i])
		}
	}
	return countUnique(obstructions)
}

func checkLoop(obstr vector) bool {
	l, d := start, vector{0, -1}

	visited := make(map[step]bool)
	visited[step{l.x, l.y, d.x, d.y}] = true

	for {
		next := add(l, d)

		if !inLab(next) {
			return false
		}

		if lab[next.y][next.x] == '#' || next == obstr {
			d = rotateRight(d)
			continue
		}

		l = next

		if visited[step{l.x, l.y, d.x, d.y}] {
			return true
		}
		visited[step{l.x, l.y, d.x, d.y}] = true
	}
}
