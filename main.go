package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var line string
	flag.StringVar(&line, "", "", "")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(0)
	} else {
		for _, word := range args {

			line += word
		}
		spl := strings.Split(line, " ")
		fmt.Println(len(spl))
	}

}
