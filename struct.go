package main

import "fmt"

type Customer struct {
	Name, Address string
	Country string
	Age int
}

func (customer Customer) order(item string) {
	fmt.Println("Customer", customer.Name, "order item", item);
}

func (customer Customer) transfer(user Customer, amount int) {
	fmt.Println("Customer", customer.Name, "transfer to", user.Name, "with amount", amount);
}

func main() {
	var customer Customer
	customer.Name = "Angga Ari Wijaya"
	customer.Address = "Surabaya, Indonesia"
	// Country and Age will be default value of data type
	fmt.Println(customer)
	fmt.Println(customer.Name)

	customer2 := Customer{
		Name: "Keenan",
		Address: "Gresik",
		Age: 4,
	}
	fmt.Println(customer2)

	customer3 := Customer{"Keenan", "Surabaya", "Indonesia", 4}
	fmt.Println(customer3)

	customer.order("Coffee");
	customer.transfer(customer2, 100);
}