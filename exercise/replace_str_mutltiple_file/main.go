package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func processFile(path, oldstr, newstr string) error {

	// Read the file

	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	// Replace the string
	output := strings.ReplaceAll(string(data), oldstr, newstr)

	return os.WriteFile(path, []byte(output), 0644)
}

func processDirectory(dir, oldstr, newstr string) error {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {

		errf := processFile(path, oldstr, newstr)
		if errf != nil {
			fmt.Println("Error processing file:", errf)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: replace <directory> <oldString> <newString>")
		return
	}

	dir := os.Args[0]
	oldstr := os.Args[1]
	newstr := os.Args[2]

	err := processDirectory(dir, oldstr, newstr)

	if err != nil {
		fmt.Println("Error processing file:", err)
	}

}
