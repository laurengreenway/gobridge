package main

import (
	"fmt"
	"strconv"
	"strings"
)

// In this exercise your task is to work with our original set of proverbs once more to count the occurrence of each character per line. You will need to make use of several of the concepts you've learned about so far, including variables, values, for and range, slices, maps and functions.

const proverbs = `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`

// Exercise 3
func main() {
	lines := strings.Split(proverbs, "\n")
	for index, line := range lines {
		letters := strings.Split(line, "")
		count := make(map[string]int)
		fmt.Printf("%d. %s\n", index+1, line)

		for _, char := range letters {
			character := strings.ToLower(char)
			count[character] = count[character] + 1
		}
		var acc string
		for idx, c := range count {
			acc += formatString(idx, c)
		}
		fmt.Println(acc)

	}
}

func formatString(char string, count int) string {
	return "'" + char + "'=" + strconv.FormatInt(int64(count), 10) + " "
}
