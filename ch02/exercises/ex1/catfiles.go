package main

import (
	"fmt"
	"os"
	"time"
)

func catFile(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	fmt.Printf("content of the file: %s\n\n%s\n", file, string(data))

	return nil
}

func main() {
	for _, file := range os.Args[1:] {
		go catFile(file)
	}

	time.Sleep(time.Minute)
}
