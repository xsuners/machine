package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len(strings.SplitN("", ".", 2)))
	fmt.Println(strings.SplitN("", ".", 2)[0])
	fmt.Println(strings.SplitN("", ".", 2)[0] == "")
}
