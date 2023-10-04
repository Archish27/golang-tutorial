package main

import "fmt"

func main() {

	//var myMap map [keyType]valueType

	var heroes map[string]string
	heroes = make(map[string]string)

	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Flash"] = "Barry Allen"

	villians["Lex Luthur"] = "Lex Luthur"

	superPets := map[int]string{
		1: "Krypto",
		2: "Bat Hound",
	}
	fmt.Printf("Batman is %v \n", heroes["Batman"])
	fmt.Println("Chip :", superPets[0])

	_, ok := superPets[3]
	if ok == false {
		fmt.Println("Pet doesn't exists")
	}
	fmt.Println("Is there a 3rd pet :", ok)

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	delete(heroes, "Flash")
}
