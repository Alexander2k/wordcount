package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string
	args := os.Args[1:]

	flag.Parse()

	if len(args) == 0 {
		fmt.Println(0)
	} else {
		line = strings.Join(args, "")
		if line == "" {
			fmt.Println(0)
		} else {
			spl := strings.Split(line, " ")
			fmt.Println(len(spl))
		}

	}

}
