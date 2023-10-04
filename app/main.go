package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("Hello ",stuff.Name)
	fmt.Println("Hello ", reflect.TypeOf(stuff.IntArrToStrArr([]int{1,2,3,4})))
}