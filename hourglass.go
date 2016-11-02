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
	var t int
	if x >= 4 || y >= 4 {
		return result
	}

	for i := 0; i < 3; i++ { // vertical dimension loop
		for j := 0; j < 3; j++ { // horizontal dimension loop
			fmt.Printf("X: %d Y: %d\n", x, y)
			t, _ = strconv.Atoi(ar[x][y])
			result = result + t
			y++
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
	offsetX := 0
	offsetY := 0
	curResult := 0
	maxSum := 0 // current maximum hourglass value
	// curResult = countHG(finalAr, offsetX, offsetY)
	// proceed glasses

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
