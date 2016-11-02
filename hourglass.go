package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countHG(ar [][]string, x int, y int) int {
	// x must be less than 6 -2 -> must be less than 4
	var result int
	if x >= 4 || y >= 4 {
		return 0
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

		}
		x++
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var arStr string
	var curResult int
	finalAr := make([][]string, 6)

	cnt := make([]byte, 256)
	reader.Read(cnt)
	arStr = string(cnt)
	fmt.Println(arStr + "\n=======")
	t1 := strings.Split(arStr, "\n")
	for i := 0; i < 6; i++ {
		fmt.Printf("Proceed string #%d: %s\n", i, t1[i])
		finalAr[i] = strings.Split(t1[i], " ")
	}

	fmt.Printf("RES: %#v", finalAr)
	offsetX := 0
	offsetY := 0
	maxSum := 0 // current maximum hourglass value
	// proceed glasses
	for i := 0; i < 6; i++ { // x-dimension loop
		for j := 0; j < 6; j++ { // y-dimension loop

		}
	}
}
