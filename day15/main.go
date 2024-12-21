package main

import (
	"bytes"
	"fmt"
	"os"
)

type vector struct {
	x, y int
}

var warehouse [][]byte
var expandedWarehouse [][]byte

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	mapBytes, movementBytes, found := bytes.Cut(b, []byte("\n\n"))
	if !found {
		panic("No empty line in input")
	}

	warehouse = bytes.Fields(mapBytes)
	expandedWarehouse = expandMap(warehouse)

	move(warehouse, movementBytes)
	fmt.Printf("Answer part 1: %d\n", sumBoxCoords(warehouse))

	move(expandedWarehouse, movementBytes)
	fmt.Printf("Answer part 2: %d\n", sumBoxCoords(expandedWarehouse))
}

func move(warehouse [][]byte, movements []byte) {
	for _, movement := range movements {
		switch movement {
		case '<':
			moveLeft(warehouse)
		case '^':
			moveUp(warehouse)
		case '>':
			moveRight(warehouse)
		case 'v':
			moveDown(warehouse)
		case '\n':
		default:
			panic(fmt.Errorf("unrecognized character %q", movement))
		}
	}
}

func moveLeft(warehouse [][]byte) {
	toCheck := []vector{findRobot(warehouse)}
	var toMove []vector

	for len(toCheck) > 0 {
		l := vector{toCheck[0].x - 1, toCheck[0].y}
		switch warehouse[l.y][l.x] {
		case '#':
			return
		case '.':
			toMove = append(toMove, toCheck[0])
		case 'O', '[', ']':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, l)
		default:
			panic(fmt.Errorf("unrecognized character %q", warehouse[l.y][l.x]))
		}
		toCheck = toCheck[1:]
	}

	for i := len(toMove) - 1; i >= 0; i-- {
		warehouse[toMove[i].y][toMove[i].x-1] = warehouse[toMove[i].y][toMove[i].x]
		warehouse[toMove[i].y][toMove[i].x] = '.'
	}
}

func moveUp(warehouse [][]byte) {
	toCheck := []vector{findRobot(warehouse)}
	var toMove []vector

	for len(toCheck) > 0 {
		u := vector{toCheck[0].x, toCheck[0].y - 1}
		switch warehouse[u.y][u.x] {
		case '#':
			return
		case '.':
			toMove = append(toMove, toCheck[0])
		case 'O':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, u)
		case ']':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, u, vector{u.x - 1, u.y})
		case '[':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, u, vector{u.x + 1, u.y})
		default:
			panic(fmt.Errorf("unrecognized character %q", warehouse[u.y][u.x]))
		}
		toCheck = toCheck[1:]
	}

	moved := make(map[vector]bool) // prevent processing same position twice
	for i := len(toMove) - 1; i >= 0; i-- {
		if !moved[toMove[i]] {
			warehouse[toMove[i].y-1][toMove[i].x] = warehouse[toMove[i].y][toMove[i].x]
			warehouse[toMove[i].y][toMove[i].x] = '.'
			moved[toMove[i]] = true
		}
	}
}

func moveRight(warehouse [][]byte) {
	toCheck := []vector{findRobot(warehouse)}
	var toMove []vector

	for len(toCheck) > 0 {
		r := vector{toCheck[0].x + 1, toCheck[0].y}
		switch warehouse[r.y][r.x] {
		case '#':
			return
		case '.':
			toMove = append(toMove, toCheck[0])
		case 'O', '[', ']':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, r)
		default:
			panic(fmt.Errorf("unrecognized character %q", warehouse[r.y][r.x]))
		}
		toCheck = toCheck[1:]
	}

	for i := len(toMove) - 1; i >= 0; i-- {
		warehouse[toMove[i].y][toMove[i].x+1] = warehouse[toMove[i].y][toMove[i].x]
		warehouse[toMove[i].y][toMove[i].x] = '.'
	}
}

func moveDown(warehouse [][]byte) {
	toCheck := []vector{findRobot(warehouse)}
	var toMove []vector

	for len(toCheck) > 0 {
		d := vector{toCheck[0].x, toCheck[0].y + 1}
		switch warehouse[d.y][d.x] {
		case '#':
			return
		case '.':
			toMove = append(toMove, toCheck[0])
		case 'O':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, d)
		case ']':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, d, vector{d.x - 1, d.y})
		case '[':
			toMove = append(toMove, toCheck[0])
			toCheck = append(toCheck, d, vector{d.x + 1, d.y})
		default:
			panic(fmt.Errorf("unrecognized character %q", warehouse[d.y][d.x]))
		}
		toCheck = toCheck[1:]
	}

	moved := make(map[vector]bool) // prevent processing same position twice
	for i := len(toMove) - 1; i >= 0; i-- {
		if !moved[toMove[i]] {
			warehouse[toMove[i].y+1][toMove[i].x] = warehouse[toMove[i].y][toMove[i].x]
			warehouse[toMove[i].y][toMove[i].x] = '.'
			moved[toMove[i]] = true
		}
	}
}

func findRobot(warehouse [][]byte) vector {
	for y, l := range warehouse {
		for x, c := range l {
			if c == '@' {
				return vector{x, y}
			}
		}
	}
	panic("robot not found")
}

func expandMap(m [][]byte) [][]byte {
	em := make([][]byte, len(m))
	for y, l := range m {
		el := make([]byte, 2*len(m[0]))
		for x, c := range l {
			switch c {
			case '#':
				el[2*x], el[2*x+1] = '#', '#'
			case 'O':
				el[2*x], el[2*x+1] = '[', ']'
			case '.':
				el[2*x], el[2*x+1] = '.', '.'
			case '@':
				el[2*x], el[2*x+1] = '@', '.'
			default:
				panic(fmt.Errorf("unknow character %q", c))
			}
		}
		em[y] = el
	}
	return em
}

func sumBoxCoords(warehouse [][]byte) int {
	ans := 0
	for y, l := range warehouse {
		for x, c := range l {
			if c == 'O' || c == '[' { // 'O' for part 1, '[' for part 2
				ans += 100*y + x
			}
		}
	}
	return ans
}
