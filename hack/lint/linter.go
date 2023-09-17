package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !isValidSnakeCase(fileName) {
			fmt.Printf("File name '%s' does not follow snake_case naming convention\n", fileName)
			os.Exit(1)
		}
	}
}

func isValidSnakeCase(s string) bool {
	snakeCasePattern := regexp.MustCompile("^[a-z0-9_]+$")
	return snakeCasePattern.MatchString(s) && strings.Contains(s, "_")
}
