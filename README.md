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

## Day 2

Conceptually simple. The implementation of part 1 was straightforward. Part 2 proved more tricky. The algorithm was simple, yet implementing it took a bit to understand better how the concept of slices is implemented and why the underlying array got modified.

Things I learned:

- [Go Slices: usage and internals](https://go.dev/blog/slices-intro).
- When using `slices.Delete` a new slices is produced, but the underlying array gets modified. Therefore, necessary to copy the data.

## Day 3

The first time this years that requires regular expressions. Conceptually the solution was straightforward for both parts. Spent most time figuring out regular expression syntax with subexpressions in Go.

Things I learned:

- Specifics of [regular expression syntax](https://pkg.go.dev/regexp/syntax).
- Using `switch` as a more terse if-then-else syntax.

## Day 4

Again not very difficult. For the first part, I doubted whether scanning the whole grid per direction (left to right, right to left, etc.) or invetigate per character. Decided for the latter.

Things I learned:

- Although I did not find it in the docs, Richard during our commute learned me you can use character literals, e.g. `'M'`.

## Day 5

Luckily saw quickly that it's simpler to evaluate the list of rules, than trying to work out if the rules are applied from the list of updates. Made use of the `Index` function from the `slices` package which removed the need for looping. First time using a struct.