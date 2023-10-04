package main

import (
	"fmt"
)

type Customer struct{
	name string
	address string
	balance float64
}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64{
	return r.length * r.height
}

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s", b.name, b.contact.fName, b.contact.lName)
}

func getCustInfo(c Customer)  {
	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
}

func newCustomerAddress(c *Customer, address string)  {
	c.address = address
}

var pl = fmt.Println

func main() {

	var tS Customer
	tS.address = "Mumbai"
	tS.name = "Tom"
	tS.balance = 10.45

	getCustInfo(tS)
	newCustomerAddress(&tS, "Chennai")
	pl("Address: ", tS.address)
	sS := Customer{"Rohan", "Pune", 100.2}
	pl("Name: ", sS.name)

	rect := rectangle{10.5, 20.5}
	pl("Area ", rect.Area())

	//Go doesnt support inheritance but support composition
	con1 := contact{
		"James", 
		"Wang",
		"1232",
	}
	business := business{
		"Business",
		"Mumbai",
		con1,
	}
	business.info()
}
