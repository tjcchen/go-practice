package main

import "fmt"

/**
 * Interface defines a group of method signatures
 */
type IF interface {
	getName() string
}

/**
 * Struct defines a group of properties
 */
type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return h.firstName + ", " + h.lastName
}

func (p Plane) getName() string {
	// The fmt.Sprintf function in the GO programming language is a function used to return a formatted string
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	interfaces := []IF{}

	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)

	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)

	for _, f := range interfaces {
		fmt.Println(f.getName()) // first, last \n benz-s
	}

	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName()) // vendor: testVendor, model: testModel
}
