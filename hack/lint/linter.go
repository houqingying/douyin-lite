package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"k8s.io/klog"
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
		// just for .go file
		//klog.Info("Checking file: ", file.Name())
		if !strings.HasSuffix(file.Name(), ".go") && !file.IsDir() {
			continue
		}
		klog.Info("Checking directory: ", file.Name())
		// if dir, recursively check
		if file.IsDir() {

			err := os.Chdir(file.Name())
			if err != nil {
				fmt.Printf("Error changing directory: %v\n", err)
				os.Exit(1)
			}
			main()
			err = os.Chdir("..")
			if err != nil {
				fmt.Printf("Error changing directory: %v\n", err)
				os.Exit(1)
			}
			continue
		}
		fileName := file.Name()
		// we don't need .go
		fileName = strings.TrimSuffix(fileName, ".go")
		if !isValidSnakeCase(fileName) {
			fmt.Printf("File name '%s' does not follow snake_case naming convention\n", fileName)
			os.Exit(1)
		}
	}
}

func isValidSnakeCase(s string) bool {
	snakeCasePattern := regexp.MustCompile("^[a-z0-9_]+$")
	return snakeCasePattern.MatchString(s)
}
