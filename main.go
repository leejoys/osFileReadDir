package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		log.Fatal("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		log.Fatal(err)
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	file, err := os.Open(path)
	defer file.Close()
	entries, err := file.ReadDir(0)
	fmt.Println(entries)
	return err
}
