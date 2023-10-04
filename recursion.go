package main

import (
	"fmt"
)

var pl = fmt.Println

func factorial(n int) int {
	if (n == 0) || (n == 1) {
		return 1
	}
	return n * factorial(n - 1)
}

func main()  {
	pl(factorial(120))
}
