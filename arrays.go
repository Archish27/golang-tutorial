package main

import(
	"fmt"
)

var pl = fmt.Println

func main() {
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1,2,3,4,5}
	pl("Index 0: ", arr2[0])
	pl("Arr length", len(arr2))
	for i := 0; i < len(arr2); i++{
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[0]); j++ {
			pl(arr3[i][j])
		}
	}

	aString := "HelloWorld"
	rArr := []rune(aString)
	for _, v := range rArr {
		fmt.Printf("Rune Array %d\n", v)
	}

	//Byte Array to String
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("Im a string", bStr)
}
