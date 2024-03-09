package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	// Open file
	file, err := os.Open(fm.InputFilePath)
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

func (fm FileManager) WriteResult(data any) error {
	// Create File
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("create fail")
	}
	encoder := json.NewEncoder(file)
	// encode data
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("fail to convert to JSON")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
