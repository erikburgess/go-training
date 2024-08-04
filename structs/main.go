package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	//	alex := person{
	//		"Alex",
	//		"Anderson",
	//	}
	//	fmt.Println(alex)
	//	fmt.Printf("%+v", alex)
	//	var alex2 person
	//
	//	fmt.Println(alex2)
	//	fmt.Printf("%+v", alex2)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipCode: 95110,
		},
	}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
