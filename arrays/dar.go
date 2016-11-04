package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	isFirstLine := true
	symbolCounter := 1 // it means first
	var typeOfQuery int
	for {

		data := make([]byte, 1)
		_, err := reader.Read(data)
		if err == io.EOF {
			break
		}

		if isFirstLine { // first line logic
			if string(data) == " " {
				continue
			} else if string(data) == "\n" {
				isFirstLine = false
				symbolCounter = 1
			} else {
				switch symbolCounter {
				case 1:
					countSeq, _ := strconv.Atoi(string(data))
					seqStore := make([][]string, countSeq)
				default:
					continue
				}
				fmt.Printf("FIRST_LINE => %s\n", string(data))
			}
		} else { // other lines logic
			if string(data) == " " {
				continue

			} else if string(data) == "\n" {
				symbolCounter = 1
			} else {
				switch symbolCounter {
				case 1: // type of query
				}
				fmt.Printf("%s", string(data))
			}
		}
		symbolCounter++
	}
}
