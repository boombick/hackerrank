package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	firstLineSlice := strings.Split(trimSuffix(firstLine, "\n"), " ")
	sliceSize, _ := strconv.Atoi(firstLineSlice[0])
	countLoop, _ := strconv.Atoi(firstLineSlice[1])

	currentIndex := sliceSize - countLoop
	finalSlice := make([]string, sliceSize)
	fmt.Printf("count loop is %d\n", countLoop)
	for i := 1; i <= sliceSize; i++ {
		line, _ := reader.ReadString(' ')
		line = strings.Trim(line, " \n")

		finalSlice[currentIndex] = line
		currentIndex++
		if currentIndex == sliceSize {
			currentIndex = 0
		}

	}
	fmt.Printf("%s\n", strings.Join(finalSlice, " "))
}
