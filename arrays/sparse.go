package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.Trim(firstLine, "\n")
	mapIndex, _ := strconv.Atoi(firstLine)

	valMap := make(map[string]int)

	for i := 1; i <= mapIndex; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.Trim(line, " \n")

		if val, ok := valMap[line]; ok {
			valMap[line] = val + 1
		} else {
			valMap[line] = 1
		}
	}

	queryLine, _ := reader.ReadString('\n')
	queryLine = strings.Trim(queryLine, "\n")
	queryIndex, _ := strconv.Atoi(queryLine)

	for i := 1; i <= queryIndex; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.Trim(line, " \n")

		if val, ok := valMap[line]; ok {
			fmt.Printf("%d\n", val)
		} else {
			fmt.Println(0)
		}
	}
}
