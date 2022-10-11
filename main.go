package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.String("", "", "")
	flag.Parse()
	args := flag.Args()
	strings.Join(args, "")
	fmt.Println(len(args))

}
