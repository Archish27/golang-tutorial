package main

import (
	"fmt"
)

type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tsp TBs) ML {
	return ML(tsp * 14.92)
}

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}


func (tsp TBs) ToMLs() ML {
	return ML(tsp * 14.79)
}

var pl = fmt.Println

func main() {

	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f ML \n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %2.f ML \n", ml2)
	pl("2 tsp + 4 tsp = ", Tsp(2) + Tsp(4))
	pl("2 tsp > 4 tsp = ", Tsp(2) > Tsp(4))
	fmt.Printf("3 tsp = %.2f ml \n", tspToML(3))
	fmt.Printf("3 tsp = %.2f ml \n", TBToML(3))

	tsp1 := Tsp(3)
	pl(tsp1.ToMLs())
}
