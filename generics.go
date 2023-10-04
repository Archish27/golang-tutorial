package main

import "fmt"

type MyConstraint interface{
	int | float64
}

//pkg.go.dev/golang.org/x/exp/constraints
func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {

	fmt.Println("5 + 4 = ", getSumGen(5,4))
	fmt.Println("5.6 + 4.7 = ", getSumGen(5.6,4.7))
		
}
