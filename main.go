package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/logo.png
var logo []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("test/files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			filename := entry.Name()

			fmt.Println(filename)

			content, _ := path.ReadFile("test/files/" + filename)
			fmt.Println(string(content))
		}
	}
}
