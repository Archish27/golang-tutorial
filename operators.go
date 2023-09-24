package main

import(
	"fmt"
)

var pl  = fmt.Println

func main()  {
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !

	iAge := 8

	if(iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday", iAge)
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday", iAge)
	} else if (iAge >= 65){
		pl("Important Birthday", iAge)
	} else {
		pl("Not Important Birthday", iAge)
	}
	pl("true", !false)
	pl("!true", !true)
}
