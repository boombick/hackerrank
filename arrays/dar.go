package dar

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		data := make([]byte, 1)
		_, err := reader.Read(data)
		if err == io.EOF {
			break
		}

		if string(data) == " " {
			continue
		}

		fmt.Printf("%s", string(data))
	}
}
