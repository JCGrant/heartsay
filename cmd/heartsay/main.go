package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JCGrant/heartsay"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	err := heartsay.Say(text)
	if err != nil {
		fmt.Println(err)
	}
}
