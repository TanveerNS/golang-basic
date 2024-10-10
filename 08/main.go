package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeToFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}

func readFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not read file: %v", err)
	}
	return string(content), nil
}

func main() {
	filename := "example.txt"

	content := "Hello, Go! This is an example of file handling.\nLet's write and read files."
	err := writeToFile(filename, content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully!")

	readContent, err := readFromFile(filename)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("\nContent of the file:")
	fmt.Println(readContent)
}
