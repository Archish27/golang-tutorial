package main

import(
	"fmt"
)

var pl = fmt.Println

func changeValue(f3 *int) {
	*f3 += 1
}

func dblArrVals(arr *[4]int){
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2 
	}
}

func getAverage(nums ...float64) float64 {
	sum := 0.0
	numSize := float64(len(nums))
	for _, value := range nums {
		sum += value
	}
	return (sum/numSize)
}

func main()  {
	f3 := 10

	f4 := 10
	var f4Ptr *int = &f4

	pl("Address: ", f4Ptr)
	pl("Value: ", *f4Ptr)
	
	pl("before: ", f3)
	//Pointer passing the address
	changeValue(&f3)
	pl("after: ", f3)

	pArr := [4]int{1,2,3,4}
	dblArrVals(&pArr)
	pl(pArr)

	iSlice := []float64{11,12,17}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))
}
