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
	for i := 0; i <= 2; i++ { // horizontal dimension loop (X)
		for j := 0; j <= 2; j++ { // vertical dimension loop (Y)
			if (j == 0 && i == 1) || (j == 2 && i == 1) {
				//				fmt.Printf("\t====> X: %d Y: %d not in operation\n", x, y)
				y++
				if y == m+3 {
					y = m
				}
				continue
			}
			t, _ = strconv.Atoi(ar[x][y])
			result = result + t
			//fmt.Printf("X: %d Y: %d M: %d VAL: %s\n", x, y, m, ar[x][y])
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
	//	fmt.Println(arStr + "\n=======")
	t1 := strings.Split(arStr, "\n")
	for i := 0; i < 6; i++ {
		//		fmt.Printf("Proceed string #%d: %s\n", i, t1[i])
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
		for i := 0; i <= 5; i++ {
			curResult = countHG(finalAr, k, i)
			//			fmt.Printf("Cur result is %d Max sum id %d\n", curResult, maxSum)
			//			fmt.Printf("==> I: %d J: %s\n", i, j)
			if curResult > maxSum {
				maxSum = curResult
			}
		}
	}
	fmt.Printf("%d", maxSum)
}

/*
-1 -1 0 -9 -2 -2
-2 -1 -6 -8 -2 -5
-1 -1 -1 -2 -3 -4
-1 -9 -2 -4 -4 -5
-7 -3 -3 -2 -9 -9
-1 -3 -1 -2 -4 -5
Incorrect input
*/
