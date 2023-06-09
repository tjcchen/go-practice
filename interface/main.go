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

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	return h.firstName + ", " + h.lastName
}

func (p Plane) getName() string {
	// The fmt.Sprintf function in the GO programming language is a function used to return a formatted string
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
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

	// struct pointer address & value
	h1 := Human{} // h1 is a struct
	p1 := &h1     // & means the pointer address, p1 is a pointer variable of h1
	v1 := *p1     // * means the pointer value
	fmt.Printf("struct: %s, pointer address: %s, pointer value: %s\n", h1, p1, v1) // { } &{ } { }

	// instantiate struct & print value
	t := MyType{Name: "test"}
	printMyType(&t)  // test

	// polymorphism
	var human IF
	human = new(Human)
	fmt.Println(human.getName()) //  , 
}

type MyType struct {
	Name string
}

func printMyType(t *MyType) {
	println(t.Name)
}
