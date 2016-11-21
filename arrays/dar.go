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

func main() {
	var line string
	var lineSlice []string
	lastAns := 0
	var querySeqPosition, x, y int
	var lenSeq int
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	firstLineSlice := strings.Split(trimSuffix(firstLine, "\n"), " ")

	seqSize, _ := strconv.Atoi(firstLineSlice[0])
	loopCounter, _ := strconv.Atoi(firstLineSlice[1])
	seqList := make([][]int, seqSize)

	for i := 1; i <= loopCounter; i++ {
		line, _ = reader.ReadString('\n')
		lineSlice = strings.Split(trimSuffix(line, "\n"), " ")
		x, _ = strconv.Atoi(lineSlice[1])
		y, _ = strconv.Atoi(lineSlice[2])
		querySeqPosition = (x ^ lastAns) % seqSize

		switch lineSlice[0] {
		case "1":
			seqList[querySeqPosition] = append(seqList[querySeqPosition], y)
		case "2":
			lenSeq = len(seqList[querySeqPosition])
			lastAns = seqList[querySeqPosition][y%lenSeq]
			fmt.Println(lastAns)
		default:
			panic("Very strange value")
		}
	}
}
