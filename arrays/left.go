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
	var s string

	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	firstLineSlice := strings.Split(trimSuffix(firstLine, "\n"), " ")
	countLoop, _ := strconv.Atoi(firstLineSlice[1])

	secondLine, _ := reader.ReadString('\n')
	secondLineSlice := strings.Split(trimSuffix(secondLine, "\n"), " ")

	for i := 1; i <= countLoop; i++ {
		s = secondLineSlice[0]
		secondLineSlice = removeIndex(secondLineSlice, 0)
		secondLineSlice = append(secondLineSlice, s)
	}
	fmt.Printf("%s\n", strings.Join(secondLineSlice, " "))
}
