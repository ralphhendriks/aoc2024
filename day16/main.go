package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

type direction int

const (
	East direction = iota
	South
	West
	North
)

type node struct {
	x, y int
	d    direction
}

type vector struct {
	x, y int
}

// Start priority queue implementation, see https://pkg.go.dev/container/heap#example-package-PriorityQueue

type Item struct {
	node     node
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority, so we use smaller than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// End priority queue implementation

var maze [][]byte

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	maze = bytes.Fields(b)
	dist, prev := shortestDistances(findStart())
	fmt.Printf("Answer part 1: %d\n", lowestScore(dist))
	fmt.Printf("Answer part 2: %d\n", tilesOnBestPath(dist, prev))
}

func findStart() node {
	for y, l := range maze {
		for x, c := range l {
			if c == 'S' {
				return node{x, y, East}
			}
		}
	}
	panic("start not found")
}

func lowestScore(distances map[node]int) int {
	x, y := findEnd()
	var found []int
	for _, d := range []direction{East, South, West, North} {
		if dist, ok := distances[node{x, y, d}]; ok {
			found = append(found, dist)
		}
	}
	return slices.Min(found)
}

func findEnd() (x int, y int) {
	for y = 0; y < len(maze); y++ {
		for x = 0; x < len(maze[y]); x++ {
			if maze[y][x] == 'E' {
				return
			}
		}
	}
	panic("end not found")
}

// Dijkstra, see https://en.wikipedia.org/wiki/Dijkstra's_algorithm#Pseudocode
func shortestDistances(source node) (map[node]int, map[node][]node) {
	distances := make(map[node]int)
	distances[source] = 0

	prev := make(map[node][]node)

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		node:     source,
		priority: 0,
		index:    0,
	}

	visited := make(map[node]bool)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		if visited[item.node] {
			continue
		}
		visited[item.node] = true

		// look ahead
		if nA := lookAhead(item.node); maze[nA.y][nA.x] != '#' {
			alt := item.priority + 1
			if currentDistance, ok := distances[nA]; !ok || alt < currentDistance {
				distances[nA] = alt
				heap.Push(&pq, &Item{node: nA, priority: alt})
				prev[nA] = []node{item.node}
			} else {
				if ok && alt == currentDistance {
					prev[nA] = append(prev[nA], item.node)
				}
			}
		}

		// turn left and right
		nL, nR := turnLeft(item.node), turnRight(item.node)
		alt := item.priority + 1000

		if currentDistance, ok := distances[nL]; !ok || alt < currentDistance {
			distances[nL] = alt
			heap.Push(&pq, &Item{node: nL, priority: alt})
			prev[nL] = []node{item.node}
		} else {
			if ok && alt == currentDistance {
				prev[nL] = append(prev[nL], item.node)
			}
		}

		if currentDistance, ok := distances[nR]; !ok || alt < currentDistance {
			distances[nR] = alt
			heap.Push(&pq, &Item{node: nR, priority: alt})
			prev[nR] = []node{item.node}
		} else {
			if ok && alt == currentDistance {
				prev[nR] = append(prev[nR], item.node)
			}
		}
	}

	return distances, prev
}

func tilesOnBestPath(distances map[node]int, prev map[node][]node) int {
	x, y := findEnd()

	var queue []node
	var minDist int
	for _, d := range []direction{East, South, West, North} {
		if dist, ok := distances[node{x, y, d}]; ok {
			if len(queue) == 0 || dist < minDist {
				queue = []node{{x, y, d}}
				minDist = dist
			} else {
				if dist == minDist {
					queue = append(queue, node{x, y, d})
				}
			}
		}
	}

	visited := make(map[vector]bool)

	for len(queue) > 0 {
		head := queue[0]
		visited[vector{head.x, head.y}] = true
		queue = append(queue, prev[head]...)
		queue = queue[1:]
	}

	return len(visited)
}

func lookAhead(n node) node {
	switch n.d {
	case East:
		return node{n.x + 1, n.y, East}
	case South:
		return node{n.x, n.y + 1, South}
	case West:
		return node{n.x - 1, n.y, West}
	case North:
		return node{n.x, n.y - 1, North}
	default:
		panic(fmt.Errorf("unknown direction %d", n.d))
	}
}

func turnLeft(n node) node {
	switch n.d {
	case East:
		return node{n.x, n.y, North}
	case South:
		return node{n.x, n.y, East}
	case West:
		return node{n.x, n.y, South}
	case North:
		return node{n.x, n.y, West}
	default:
		panic(fmt.Errorf("unknown direction %d", n.d))
	}
}

func turnRight(n node) node {
	switch n.d {
	case East:
		return node{n.x, n.y, South}
	case South:
		return node{n.x, n.y, West}
	case West:
		return node{n.x, n.y, North}
	case North:
		return node{n.x, n.y, East}
	default:
		panic(fmt.Errorf("unknown direction %d", n.d))
	}
}
