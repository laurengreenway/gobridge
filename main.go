package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Exercise 4

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "/proverbs.txt")
	check(err)
	proverbs := string(data)

	lines := strings.Split(proverbs, "\n")
	for index, line := range lines {
		letters := strings.Split(line, "")
		count := make(map[string]int)
		fmt.Printf("%d. %s\n", index+1, line)

		for _, char := range letters {
			character := strings.ToLower(char)
			if _, ok := count[char]; ok {
				val := count[character]
				count[character] = val + 1
			} else {
				count[character] = 1
			}
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

// Instructor solution
func charCount(line string) map[rune]int {
	// runes are int32 type, golang is a unicode friendly language so characters are associated with a numerical map (?)
	m := make(map[rune]int, 0)
	for _, char := range line {
		// if no value there is a zero value, therefore we do not need to do a check for the keys existance!
		m[char] = m[char] + 1
	}
	return m
}

// fmt.Println() %c refers to the the unicode reference
