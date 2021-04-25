package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const DIRECTORY = "/files/"

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files := listFiles(path + DIRECTORY)

	for _, file := range *files {
		fmt.Println(file)

		fmt.Println()
		err = readFile(file, processLine)
		fmt.Println()

		if err != nil {
			fmt.Println("Error")
		}
	}
}

func listFiles(path string) *[]string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	return &files
}

func readFile(filePath string, handleLine func(string)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() error {
		if err = file.Close(); err != nil {
			return err
		}
		return nil
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		handleLine(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}

func processLine(line string) {
	fmt.Println(line)
}
