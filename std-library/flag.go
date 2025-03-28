package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "angga.ari", "Database username")
	password := flag.String("password", "secret", "Database password")
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 3306, "Database port")
	secure := flag.Bool("secure", false, "Database secure")

	flag.Parse()

	// go run flag.go --username=root --password="my password" --secure=true
	fmt.Println(*username)
	fmt.Println(*password)
	fmt.Println(*host)
	fmt.Println(*port)
	fmt.Println(*secure)
}