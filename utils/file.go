package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) ([]string, error) {
	ret := []string{}

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return ret, nil
	}

	return ret, nil
}
