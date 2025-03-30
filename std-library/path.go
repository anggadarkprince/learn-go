package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	// using slash
	fmt.Println(path.Dir("src/images/file.jpg"))
	fmt.Println(path.Base("images/2025/10/01/file.jpg"))
	fmt.Println(path.Ext("images/2025/10/01/file.jpg"))
	fmt.Println(path.Split("images/2025/10/01/file.jpg"))
	fmt.Println(path.Join("images", "2025", "10", "01", "file.jpg"))

	// depends on operating system (unix or windws)
	fmt.Println(filepath.Dir("src/images/file.jpg"))
	fmt.Println(filepath.Base("images/2025/10/01/file.jpg"))
	fmt.Println(filepath.Ext("images/2025/10/01/file.jpg"))
	fmt.Println(filepath.IsAbs("/Users/angga/images/2025/10/01/file.jpg"))
	fmt.Println(filepath.IsLocal("images/2025/10/01/file.jpg"))
	fmt.Println(filepath.Split("images/2025/10/01/file.jpg"))
	fmt.Println(filepath.Join("images", "2025", "10", "01", "file.jpg"))
}