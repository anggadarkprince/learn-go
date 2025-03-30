package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)
	for {
		line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(prefix)
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hey\n")
	_, _ = writer.WriteString("my name Angga\n")
	writer.Flush()
}