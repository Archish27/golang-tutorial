package main

import(
	"fmt"
)

var pl = fmt.Println

func sayHello() {
	pl("Hello World")
}

func getSum(x int, y int) (int,int) {
	return x + y, x - y
}

func getQuotient(x float64, y float64) (ans float64, err error)  {
	if( y == 0){
		return -1, fmt.Errorf("Zero divide error")
	}
	return x / y, nil
}

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main()  {
	// func funcName(parameters) returnType {BODY}
	sayHello()
	add, diff := getSum(10, 20)
	pl("Res ", add, diff)
	ans, err := getQuotient(1, 0)
	pl("Res ", ans, err)

	ans1, err1 := getQuotient(1, 10)
	pl("Res ", ans1, err1)

	pl(getSum2(1,2,3,4,5,6))

	//pass by value
	pl(getArraySum([]int{1,2,3,4}))

}
