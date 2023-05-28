package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed hohpe.png
var logo []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("hohpe_next.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

}
