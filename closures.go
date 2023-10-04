package main

import (
	"fmt"
)

var pl = fmt.Println


func useFunc(f func(int, int) int, x, y int)  {
	pl("Answer ", (f(x,y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main()  {
	intSum := func (x, y int) int {
		return x + y
	}

	pl("5 + 4 = ", intSum(5, 4))

	//Closure can change value in function
	sampl1 := 1
	changeVar := func() {
		sampl1 += 1 
	}
	changeVar()
	pl("sampl1 = ", sampl1)

	useFunc(sumVals, 5, 8)
}
