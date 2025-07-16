package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func grep(match string, file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	if strings.Contains(string(data), match) {
		fmt.Printf("the file %s contains the string \"%s\"\n", file, match)
	}

	return nil
}

func main() {
	match := os.Args[1]

	entries, err := os.ReadDir(os.Args[2])

	if err != nil {
		fmt.Println("error to read directory", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			go grep(match, fmt.Sprintf("%s/%s", os.Args[2], entry.Name()))
		}
	}

	time.Sleep(time.Minute)
}
