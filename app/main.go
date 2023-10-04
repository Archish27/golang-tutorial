package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)

func main()  {
	fmt.Println("Hello ",stuff.Name)
	fmt.Println("Hello ", reflect.TypeOf(stuff.IntArrToStrArr([]int{1,2,3,4})))

	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(1)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(2010)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day : %d/%d/%d", date.Day(), date.Month(), date.Year())
}