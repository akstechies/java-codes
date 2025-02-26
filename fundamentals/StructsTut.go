package main

import "fmt"

type customer struct {
	name  string
	age   int
	place string
}

func getCustInfo(c customer) {
	fmt.Printf("%s age is %d and he lives in %s", c.name, c.age, c.place)
}

func updateCustPlace(c *customer, place string) {
	c.place = place
}

// Calculate Rectangel Area using methods

type rectangle struct {
	length, breadth float64
}

// you will see here that we will define paranthesis before funcName so that we can call the method using .funcName

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}

func StructsTut() {
	cust1 := customer{"Ayush", 29, "Lucknow"}
	fmt.Println(cust1)

	var cust2 customer
	cust2.name = "Nayan"
	cust2.age = 24
	cust2.place = "India"
	fmt.Println(cust2)

	cust3 := customer{name: "Anand", age: 64, place: "Mumbai"}
	fmt.Println(cust3)

	updateCustPlace(&cust1, "Bharat")
	getCustInfo(cust1)
	fmt.Println("")
	// Now calc rect area
	rect1 := rectangle{length: 33, breadth: 45}
	areaOfRect := rect1.Area()
	fmt.Printf("Area of rect is %.2f", areaOfRect)

	// or we can separate it using variable for contact then assign it in buisness
	biz1 := buisness{name: "AKS Techies", year: 2019, contact: contact{
		fName: "Ayush", lName: "Srivastava", age: 29,
	},
	}
	biz1.bizInfo()
}

// NESTED STRUCTS
type contact struct {
	fName, lName string
	age          int
}

type buisness struct {
	name    string
	year    int
	contact // or can also do ~ anyVar contacts
}

func (b buisness) bizInfo() {
	fmt.Printf("\n%s bizness owner contact are %s", b.name, b.contact.fName)
	fmt.Println("Business - ", b)
}
