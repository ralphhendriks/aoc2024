# Advent of Code 2024

My attempts at solving the programming puzzles of [Advent of Code 2024](https://adventofcode.com/2024). This year, I'm using the advent calendar to learn a bit of [Go](https://go.dev/).

# Useful Go learning resources

- [Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)

# Notes per day

## Day 1

Not a very difficult assignment. Great to dip my toes in the water about slices and maps in Go. Need to study a bit more on dynamic sizing vs. a priori known lenghts.

Things I learned:

- The [`Fields`](https://pkg.go.dev/strings#Fields) function in the strings package is handy for splitting strings on arbitrary whitespace.
- Rewinding a file can be done with `Seek(0, io.SeekStart)`.