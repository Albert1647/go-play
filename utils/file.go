package utils

import (
	"bufio"
	"errors"
	"os"
)

func GetStringsFromFile(path string) ([]string, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("HUH WHERE IS THE FILE???")
	}
	// Scan text
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Scan Fail
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("?? WHAT IS THIS LANGUAGE")
	}
	file.Close()
	return lines, nil
}

func WriteJSON(data interface{}) {

}
