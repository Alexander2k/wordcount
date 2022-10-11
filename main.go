package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	s := strings.Split(text, " ")
	fmt.Println(len(s))
}
