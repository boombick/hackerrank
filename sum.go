package main // Read from standard input
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cntStr string
	var arItem string
	curItem := 0
	t := 0

	for {
		cnt := make([]byte, 1)
		reader.Read(cnt)
		if string(cnt) == "\n" {
			break
		} else {
			cntStr = cntStr + string(cnt)
		}
	}

	//arIndex, _ := strconv.Atoi(cntStr)

	for {

		data := make([]byte, 1)
		_, err := reader.Read(data)
		//fmt.Printf("Data: %s\n", data)
		if string(data) == " " || err == io.EOF {
			t, _ = strconv.Atoi(arItem)
			curItem = curItem + t
			t = 0
			arItem = ""
			if err == io.EOF {
				break
			}
		} else {
			arItem = arItem + string(data)
		}
	}

	fmt.Println(curItem)
}
