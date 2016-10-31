package main // Read from standard input
import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Read all data from stdin, passing the data as parts into the channel
	// for processing.
	parts := 0
	//isSpace := false
	var cntStr string
	//var arItem string

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
	var finArray []int
	finArray = make([]int, arIndex)

	for {
		data := make([]byte, 1)
		_, err := reader.Read(data)

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Problems reading from input: %s", err)
		}
		parts++

		//if data == " "
	}
	log.Print(finArray)
	log.Print(finArray)
}
