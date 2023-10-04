package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	args := os.Args[1:]
	var iArgs = []int{}
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	pl("Max: ", max)
}
