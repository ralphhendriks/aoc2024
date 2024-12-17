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

## Day 6

Part 1 was quickly done. Part 2 took way more time than I'd like to admit. Finally, I figured out that I overlooked that multiple right turns can be made before advancing to the next position. I found this out only when looking for a hint and a Redditor suggested the following minimalistic test case:

```
###
..#
^##
```

## Day 7

The first time with big numbers, so the first time using `int64`. The `strconv.ParseInt` method is a bit more cumbersome than `strconv.Atoi`, but that's only a minor annoyance. I implemented the concat operator for part 2 by going back to string representation, which feels a bit hacky, but it was the fastest day of getting it done as I did not find an int-only log10 implementation in the standard library.

## Day 8

Not difficult, but a bit quick and dirty solution. First time passing a function as an argument to swap out the antinode pattern/algorithm.

## Day 9

Conceptually, again, it was clear quite quickly how to approach. The second part took me a bit longer, as I initially did not realize that it is not necessary to add the moved files to the free blocks, as files with a lower ID cannot be moved there anyway. This simplified the approach considerably.

## Day 10

Again, this wasn't a too difficult puzzle. The algorithm was immediately clear and it was a good opportunity to test a recursive algorithm in Go. This worked out to be a lucky choice, as the implementation already calculated all the trails needed for part 2. Allowed me to experiment passing an anonymous function just for the sake of it.

Things I learned: 

- When appending two slices, the syntax `append(slice, anotherSlice...)` must literally include the ellipsis `...`. This is a special syntax for [variadic functions](https://gobyexample.com/variadic-functions) if there are already multiple arguments in a slice.

## Day 11

This took a while to figure it out. Tried a recursive approach, also with a cache/memoization on the side, but it blew up. Finally an iterative solution while keeping a histogram of the stone numbers proved a performant and simple option.

## Day 12

This was the most challenging and the most fun puzzle so far, also resulting in the longest Go program.

Things I learned:

- The algorithm needed to solve this is related to an area of computer science called [Connected-component labeling (CCL)](https://en.wikipedia.org/wiki/Connected-component_labeling).
- There are 1-pass and 2-pass algorithms described for this class of problems. I spent some time reading about the _H_oshen-Kopelman algorithm_ (see [here](https://en.wikipedia.org/wiki/Hoshen%E2%80%93Kopelman_algorithm) and [here](https://www.ocf.berkeley.edu/~fricke/projects/hoshenkopelman/hoshenkopelman.html)). Eventually, I opted for a simpler approach with a 1-pass algorithm.

## Day 13

This was an interesting exercise using some linear algebra. It made me refresh my knowledge about e.g. Cramer's rule.

Things I learned:

- Basic linear algebra refresher.
- I somehow missed that the shorthand notation for defining variables  of the same type also works for function parameters, e.g. `func foo(a, b int)`.
