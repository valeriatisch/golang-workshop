package main

import (
	"fmt"
	"github.com/valeriatisch/golang-workshop/02_packages/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello"))
	fmt.Println(cmp.Diff("Hello", "Hello"))
}
