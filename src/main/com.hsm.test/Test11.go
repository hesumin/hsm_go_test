package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkDir(dir string, indent string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		fmt.Println(indent + entry.Name())
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), indent+"  ")
		}
	}
}

func main() {
	walkDir(".", "")
}
