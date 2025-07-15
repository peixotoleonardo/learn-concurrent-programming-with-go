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
		fmt.Printf("the file %s contains a the string \"%s\"\n", file, match)
	}

	return nil
}

func main() {
	files := os.Args[2:]

	for _, file := range files {
		go grep(os.Args[1], file)
	}

	time.Sleep(time.Minute)
}
