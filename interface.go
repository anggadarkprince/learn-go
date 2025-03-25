package main

import "fmt"

type HasName interface {
	GetName() string
}

// implementation in Customer
type Customer struct {
	Name string
}

func (customer Customer) GetName() string {
	return customer.Name
}

// implementation in Admin
type Admin struct {
	Title string
}

func (admin Admin) GetName() string {
	return admin.Title
}

// Method that access interface
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Interface empty or "any" to follow contract with any value
func debugData(data any) interface{} {
	return data
}

func main() {
	customer := Customer{Name: "Angga"}
	SayHello(customer)

	admin := Admin{Title: "Administrtaor"}
	SayHello(admin)

	fmt.Println(debugData("Hello"))
	fmt.Println(debugData(12))
	fmt.Println(debugData(false));
}