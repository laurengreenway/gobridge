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
