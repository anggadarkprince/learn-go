package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// go lang embed is a feature that allows you to include files and directories in your Go binary at compile time.
// This is useful for embedding static assets like HTML, CSS, JavaScript, images, etc. into your Go applications.
// The embed package provides a way to access these embedded files at runtime.
// It is available in Go 1.16 and later versions.
// The embed package allows you to embed files and directories into your Go binary.
// You can use the //go:embed directive to specify which files or directories to embed.
// The embedded files can be accessed using the embed.FS type, which provides methods to read and manipulate the embedded files.


//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed file.png
var file []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("file_copy.png", file, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}

//go:embed files/file1.txt
//go:embed files/file2.txt
//go:embed files/file3.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	data, err := files.ReadFile("files/file1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data, err = files.ReadFile("files/file2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data, err = files.ReadFile("files/file3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

//go:embed files/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		data, err := path.ReadFile("files/" + entry.Name())
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}