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

// Left shift arrays -- hackerrank
func main() {
	var s, d []string

	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	firstLineSlice := strings.Split(trimSuffix(firstLine, "\n"), " ")
	arrayLen, _ := strconv.Atoi(firstLineSlice[0])
	countLoop, _ := strconv.Atoi(firstLineSlice[1])

	secondLine, _ := reader.ReadString('\n')
	secondLineSlice := strings.Split(trimSuffix(secondLine, "\n"), " ")
	fmt.Printf("%#v\n%#v\n%#v", arrayLen, countLoop, secondLineSlice)

	for i := 1; i <= countLoop; i++ {
	}
}
