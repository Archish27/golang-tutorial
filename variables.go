package main

import(
	"fmt"
	"reflect"
	"strconv"
)

var pl  = fmt.Println

func main()  {
	//camel case naming convention
	// var vName string = "Derek"

	// var v1,v2 = 1.2, 3.4

	// var v3 = "Hello"

	// v4 := 2.4

	// v4 = 5.4

	// Data types 
	//               int, float64, bool, string, rune
	// Default Type    0     0.0   false   ""

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.4))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf(false))
	pl(reflect.TypeOf('😋')) //int32
	pl(reflect.TypeOf("😋")) //string

	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6, reflect.TypeOf(cV6))

	cV7 := "5000.1"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
	
}
