package main

import(
	"fmt"
	"time"
	"math/rand"
	"math"
)

var pl  = fmt.Println

func main() {
	pl("5 + 4 = ", 5 + 4)
	pl("5 - 4 = ", 5 - 4)
	pl("5 * 4 = ", 5 * 4)
	pl("5 / 4 = ", 5 / 4)
	pl("5 ^ 4 = ", 5 ^ 4)
	pl("5 % 4 = ", 5 % 4)
	pl("5 | 4 = ", 5 | 4)
	pl("5 & 4 = ", 5 & 4)
	mInt := 1
	mInt++
	mInt += 1
	pl(mInt)

	pl("Float Precision = ", 0.1111111111111111111111111111111111 + 0.1111111111111111111111111111111111)
	
	//1/1/70
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random: ", randNum)

	pl("Abs(-10)", math.Abs(-10))
	pl("Pow(4, 2)", math.Pow(4, 2))
	pl("Sqrt(16)", math.Sqrt(16))
	pl("Ceil(4.2)", math.Ceil(4.2))
	pl("Floor(4.2)", math.Floor(4.2))
	pl("Round(4.2)", math.Round(4.2))
	pl("Min(4, 2)", math.Min(4,2))
	pl("Max(4, 2)", math.Max(4,2))

	r90 := 90 * math.Pi / 180
	d90 := r90 * ( 180 / math.Pi )
	fmt.Printf("%.2f radius = %.2f degree\n", r90, d90)
}
