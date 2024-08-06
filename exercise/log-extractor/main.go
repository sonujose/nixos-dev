package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open log file
	file, err := os.Open("logfile.log")
	if err != nil {
		fmt.Println("Error fetching log file")
	}

	defer file.Close()

	// Define a regex pattern to extract specific information
	pattern := regexp.MustCompile(`ERROR: (.*)`)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		match := pattern.FindStringSubmatch(line)

		if len(match) > 0 {
			fmt.Println("Extracted log", match[1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
