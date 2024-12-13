package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return []string{}, nil
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return []string{}, err
	}
	fmt.Print(fileBytes)

	return strings.Split(string(fileBytes), "\n"), nil
}
