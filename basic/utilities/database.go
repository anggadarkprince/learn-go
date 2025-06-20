package utilities

import "fmt"

var connection string

// auto called when it's imported
func init() {
	fmt.Println("Conneting to database...")
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}