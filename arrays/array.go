package main // Read from standard input
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cntStr string
	var arItem string
	var curIndex int

	parts := 0

	for {
		cnt := make([]byte, 1)
		reader.Read(cnt)
		if string(cnt) == "\n" {
			break
		} else {
			cntStr = cntStr + string(cnt)
		}
	}

	arIndex, _ := strconv.Atoi(cntStr)
	var finArray []string
	finArray = make([]string, arIndex)

	for {

		data := make([]byte, 1)
		_, err := reader.Read(data)

		if string(data) == " " || err == io.EOF {
			parts++ // increment counter
			// make integer and put it into end of slice
			curIndex = arIndex - parts
			finArray[curIndex] = arItem
			arItem = ""
			if err == io.EOF {
				break
			}
		} else {
			arItem = arItem + string(data)
		}
	}
	fmt.Println(strings.Join(finArray, " "))
}
