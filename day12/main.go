package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type vector struct {
	x, y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var garden []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		garden = append(garden, scanner.Text())
	}

	regions := findRegions(garden)
	ans1, ans2 := 0, 0
	for _, r := range regions {
		ans1 += len(r) * perimeter(r, garden)
		ans2 += len(r) * sides(r, garden)
	}
	fmt.Printf("Answer part 1: %d\n", ans1)
	fmt.Printf("Answer part 2: %d\n", ans2)
}

func findRegions(g []string) [][]vector {
	visited := make(map[vector]bool)
	var regions [][]vector
	for y, row := range g {
		for x := range row {
			if visited[vector{x, y}] { // plot already visited, continue
				continue
			}
			// plot hasn't been visited
			var region []vector
			queue := []vector{(vector{x, y})}
			for len(queue) > 0 {
				if visited[queue[0]] {
					// plot already visited, pop the queue and continue
					queue = queue[1:]
					continue
				}
				if n := north(queue[0]); inGarden(n, g) && plant(n, g) == plant(queue[0], g) {
					// the plot on the north is in the garden and of the same type of plant
					queue = append(queue, n)
				}
				if s := south(queue[0]); inGarden(s, g) && plant(s, g) == plant(queue[0], g) {
					// the plot on the south is in the garden and of the same type of plant
					queue = append(queue, s)
				}
				if e := east(queue[0]); inGarden(e, g) && plant(e, g) == plant(queue[0], g) {
					// the plot on the east is in the garden and of the same type of plant
					queue = append(queue, e)
				}
				if w := west(queue[0]); inGarden(w, g) && plant(w, g) == plant(queue[0], g) {
					// the plot on the west is in the garden and of the same type of plant
					queue = append(queue, w)
				}
				// add the current plot to the region
				region = append(region, queue[0])
				// mark the current plot as visited
				visited[queue[0]] = true
				// pop the queue
				queue = queue[1:]
			}
			regions = append(regions, region)
		}
	}
	return regions
}

func north(p vector) vector {
	return vector{p.x, p.y - 1}
}

func south(p vector) vector {
	return vector{p.x, p.y + 1}
}

func east(p vector) vector {
	return vector{p.x + 1, p.y}
}

func west(p vector) vector {
	return vector{p.x - 1, p.y}
}

func inGarden(p vector, g []string) bool {
	return p.x >= 0 && p.x < len(g[0]) && p.y >= 0 && p.y < len(g)
}

func plant(p vector, g []string) byte {
	return g[p.y][p.x]
}

func perimeter(region []vector, g []string) int {
	ans := 0
	for _, p := range region {

		if n := north(p); !inGarden(n, g) || plant(n, g) != plant(p, g) {
			ans += 1
		}
		if s := south(p); !inGarden(s, g) || plant(s, g) != plant(p, g) {
			ans += 1
		}
		if e := east(p); !inGarden(e, g) || plant(e, g) != plant(p, g) {
			ans += 1
		}
		if w := west(p); !inGarden(w, g) || plant(w, g) != plant(p, g) {
			ans += 1
		}
	}
	return ans
}

func sides(region []vector, g []string) int {
	minX, maxX, minY, maxY := region[0].x, region[0].x, region[0].y, region[0].y
	for _, r := range region {
		if r.x < minX {
			minX = r.x
		}
		if r.x > maxX {
			maxX = r.x
		}
		if r.y < minY {
			minY = r.y
		}
		if r.y > maxY {
			maxY = r.y
		}
	}

	ans := 0

	// scan west to east, then north to south
	for y := minY; y <= maxY; y++ {
		stateN, stateS := false, false
		for x := minX; x <= maxX; x++ {
			p := vector{x, y}

			if slices.Index(region, p) < 0 { // not part of region
				stateN, stateS = false, false
				continue
			}

			if n := north(p); inGarden(n, g) && plant(n, g) == plant(p, g) {
				stateN = false
			} else {
				if !stateN {
					ans += 1
					stateN = true
				}
			}

			if s := south(p); inGarden(s, g) && plant(s, g) == plant(p, g) {
				stateS = false
			} else {
				if !stateS {
					ans += 1
					stateS = true
				}
			}
		}
	}

	// scan north to south, then west to east
	for x := minX; x <= maxX; x++ {
		stateW, stateE := false, false
		for y := minY; y <= maxY; y++ {
			p := vector{x, y}

			if slices.Index(region, p) < 0 { // not part of region
				stateW, stateE = false, false
				continue
			}

			if w := west(p); inGarden(w, g) && plant(w, g) == plant(p, g) {
				stateW = false
			} else {
				if !stateW {
					ans += 1
					stateW = true
				}
			}

			if e := east(p); inGarden(e, g) && plant(e, g) == plant(p, g) {
				stateE = false
			} else {
				if !stateE {
					ans += 1
					stateE = true
				}
			}
		}
	}

	return ans
}
