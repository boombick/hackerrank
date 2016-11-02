package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countHG(ar [][]string, x int, y int) int {
	// x must be less than 6 -2 -> must be less than 4
	result := 0
	var t, m int
	if x >= 4 || y >= 4 {
		return result
	}
	m = y
	for i := 0; i < 3; i++ { // vertical dimension loop
		for j := 0; j < 3; j++ { // horizontal dimension loop
			fmt.Printf("X: %d Y: %d M: %d VAL: %s\n", x, y, m, ar[x][y])
			if j == 1 {
				t, _ = strconv.Atoi(ar[x][y])
				result = result + t
			}
			y++
			if y == m+3 {
				y = m
			}
		}
		x++
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var arStr string
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

	//fmt.Printf("RES: %#v", finalAr)
	//offsetX := 0
	//offsetY := 0
	curResult := 0
	maxSum := 0 // current maximum hourglass value
	// curResult = countHG(finalAr, offsetX, offsetY)
	// proceed glasses
	for k := 0; k < 6; k++ {
		for i, j := range finalAr[k] {
			curResult = countHG(finalAr, k, i)
			fmt.Printf("Cur result is %d Max sum id %d\n", curResult, maxSum)
			fmt.Printf("==> I: %d J: %s\n", i, j)
			if curResult > maxSum {
				maxSum = curResult
			}
		}
	}
	fmt.Printf("Result is: %d", maxSum)
}

/*
echo -ne "1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 -9 0
0 0 -1 0 0 0
0 0 0 0 -2 0
0 0 0 0 0 4
" | ./bin/hourglass
*/
