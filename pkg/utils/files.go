package utils

import (
	"bytes"
	"fmt"
	"os"
)

func GetFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	return file, nil
}

func SaveFile(folder, fileName string, data bytes.Buffer) error {
	filePath := fmt.Sprintf("%s/%s", folder, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}

	_, err = data.WriteTo(file)
	if err != nil {
		return fmt.Errorf("cannot write data to file: %w", err)
	}

	return nil
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("cannot remove file: %w", err)
	}

	return nil
}
