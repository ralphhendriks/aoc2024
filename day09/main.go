package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)

	fmt.Printf("Answer part 1: %d\n", part1(input))
	fmt.Printf("Answer part 2: %d\n", part2(input))
}

func part1(input string) int {
	var mem []int

	// write input to slice
	for i, r := range string(input) {
		var id int
		if i%2 == 1 { // mem is free
			id = -1
		} else {
			id = i / 2 // mem is used
		}
		for k := 0; k < int(r-'0'); k++ {
			mem = append(mem, id)
		}
	}

	// defrag by block
	for i, j := 0, len(mem)-1; i < j; { // i points to the first element, j points to the last element
		if mem[i] > -1 {
			i++
			continue
		}
		if mem[j] < 0 {
			j--
			continue
		}
		mem[i] = mem[j]
		mem[j] = -1
	}

	return checksum(mem)
}

func checksum(mem []int) int {
	ans := 0
	for i, id := range mem {
		if id > -1 {
			ans += i * id
		}
	}
	return ans
}

type memSlice struct {
	id, start, length int
}

func part2(input string) int {
	var files []memSlice // slice of files
	var free []memSlice  // slice of free blocks

	// write input to slices used and free
	k := 0
	for i, r := range string(input) {
		size := int(r - '0')
		if i%2 == 1 { // mem is free
			free = append(free, memSlice{-1, k, size})
		} else { // mem is used
			files = append(files, memSlice{i / 2, k, size})
		}
		k += size
	}

	// defrag by file
	for i := len(files) - 1; i >= 0; i-- { // loop through files with decreasing id
		if j := slices.IndexFunc(free, func(m memSlice) bool {
			return m.length >= files[i].length && m.start < files[i].start
		}); j > -1 {
			files[i] = memSlice{files[i].id, free[j].start, files[i].length}
			free[j] = memSlice{-1, free[j].start + files[i].length, free[j].length - files[i].length}
		}
	}

	return (checksum2(files))
}

func checksum2(files []memSlice) int {
	ans := 0
	for _, f := range files {
		for k := f.start; k < f.start+f.length; k++ {
			ans += f.id * k
		}
	}
	return ans
}
