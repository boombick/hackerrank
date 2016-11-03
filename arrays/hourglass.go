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
		fmt.Printf("Fast return: x = %d, y = %d", x, y)
		return result
	}
	m = y
	for i := 0; i <= 2; i++ { // horizontal dimension loop (X)
		for j := 0; j <= 2; j++ { // vertical dimension loop (Y)
			if (j == 0 && i == 1) || (j == 2 && i == 1) {
				//fmt.Printf("\t====> X: %d Y: %d not in operation\n", x, y)
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
			fmt.Printf("%d ", t)
		}
		x++
		fmt.Println("")
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

	fmt.Printf("%#v", finalAr)
	fmt.Println("\n========")
	//fmt.Printf("RES: %#v", finalAr)
	//offsetX := 0
	//offsetY := 0
	curResult := 0
	maxSum := -100 // current maximum hourglass value
	itercount := 0
	// curResult = countHG(finalAr, offsetX, offsetY)
	// proceed glasses
	for k := 0; k <= 3; k++ {
		for i := 0; i <= 3; i++ {
			itercount++
			curResult = countHG(finalAr, k, i)
			fmt.Printf("===> Cur result is %d Max sum id %d\n", curResult, maxSum)
			if curResult > maxSum {
				maxSum = curResult
			}
			fmt.Printf("============================[ ITERATION #%d COMPLETE; MAX RESULT IS %d ]=======================================\n\n", itercount, maxSum)
		}
	}
	fmt.Printf("%d", maxSum)
}

/*
3 7 -3 0 1 8
1 -4 -7 -8 -6 5
-8 1 3 3 5 7
-2 4 3 1 2 7
2 4 -5 1 8 4
5 -7 6 5 2 8
Incorrect output (expected output 33)
*/
