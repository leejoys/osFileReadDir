package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	log.Fatal("usage go run main.go . [-f]")
	// }
	// path := os.Args[1]
	// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	path := "./testdata"
	printFiles := false
	err := dirTree(out, path, printFiles)
	if err != nil {
		log.Fatal(err)
	}
}

//todo https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/path/filepath/path.go;l=400

func dirTree(out *os.File, path string, printFiles bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	files, err := file.ReadDir(0)
	if err != nil {
		return err
	}
	var result string
	var space string
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return err
		}
		if info.IsDir() {
			result = result + fmt.Sprintf("%s└───%s\n", space, info.Name())
			newPath := path + fmt.Sprintf("/%s", info.Name())
			space = space + "\t"
			err := dirTree(out, newPath, printFiles)
			if err != nil {
				return err
			}
			continue
		}
		result = result + fmt.Sprintf("%s├───%s (%db)\n", space, info.Name(), info.Size())

	}
	fmt.Fprint(out, result)
	return nil
}
