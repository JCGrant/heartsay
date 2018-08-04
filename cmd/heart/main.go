package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JCGrant/heart"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	err := heart.Say(text)
	if err != nil {
		fmt.Println(err)
	}
}
