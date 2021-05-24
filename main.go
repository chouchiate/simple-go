package main

import (
	"fmt"
	"example.com/simple-go/morestrings"
	"github.com/google/go-cmp/cmp"

)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!321oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World","Hello Go"))

}