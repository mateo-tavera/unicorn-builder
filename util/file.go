package util

import (
	"bufio"
	"fmt"
	"os"
)

// Reads values from specified file
func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Please try later, unicorn factory unavailable")
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
