package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	name := "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(name))
}
