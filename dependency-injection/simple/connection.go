package simple

import "fmt"

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Closing connection to file:", c.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{
		File: file,
	}
	return connection, func() {
		connection.Close()
	}
}