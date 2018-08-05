package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/JCGrant/heartsay"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buffer bytes.Buffer
	buffer.ReadFrom(reader)
	text := buffer.String()
	text = strings.Trim(text, "\n")
	err := heartsay.Say(text)
	if err != nil {
		fmt.Println(err)
	}
}
