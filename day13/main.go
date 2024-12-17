package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type clawMachine struct {
	aX, aY, bX, bY, pX, pY int
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	clawMachines := parseMachines(b)

	fmt.Printf("Answer part 1: %d\n", part1(clawMachines))
	fmt.Printf("Answer part 2: %d\n", part2(clawMachines))
}

func part1(clawMachines []clawMachine) int {
	ans := 0
	for _, cm := range clawMachines {
		a, b := solveMachine(cm)
		if a >= 0 && b >= 0 && a <= 100 && b <= 100 {
			ans += 3*a + b
		}
	}
	return ans
}

func part2(clawMachines []clawMachine) int {
	const offset int = 10000000000000
	ans := 0
	for _, cm := range clawMachines {
		a, b := solveMachine(clawMachine{cm.aX, cm.aY, cm.bX, cm.bY, cm.pX + offset, cm.pY + offset})
		if a >= 0 && b >= 0 {
			ans += 3*a + b
		}
	}
	return ans
}

func parseMachines(b []byte) []clawMachine {
	re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)
	var machines []clawMachine
	for _, match := range re.FindAllSubmatch(b, -1) {
		var cm clawMachine
		for i, s := range match[1:] {
			num, err := strconv.Atoi(string(s))
			if err != nil {
				panic(err)
			}
			switch i {
			case 0:
				cm.aX = num
			case 1:
				cm.aY = num
			case 2:
				cm.bX = num
			case 3:
				cm.bY = num
			case 4:
				cm.pX = num
			case 5:
				cm.pY = num
			}
		}
		machines = append(machines, cm)
	}
	return machines
}

func solveMachine(cm clawMachine) (int, int) {
	det := cm.aX*cm.bY - cm.aY*cm.bX
	if det == 0 {
		if cm.aX*cm.pY == cm.aY*cm.pX && cm.bX*cm.pY == cm.bY*cm.pX {
			// infinitely many solutions exist, pressing B is cheapest
			return 0, cm.pX / cm.bX
		} else {
			// no solution exists
			return -1, -1
		}
	}

	// A unique solution exists. Use Cramer's rule for a direct solution
	detX := cm.pX*cm.bY - cm.pY*cm.bX
	detY := cm.aX*cm.pY - cm.aY*cm.pX
	if detX%det != 0 || detY%det != 0 {
		// no integer solution exists
		return -1, -1
	}
	return detX / det, detY / det
}
